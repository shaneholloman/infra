package handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"sync"
	"syscall"

	"github.com/e2b-dev/infra/packages/envd/internal/logs"
	"github.com/e2b-dev/infra/packages/envd/internal/permissions"
	rpc "github.com/e2b-dev/infra/packages/envd/internal/services/spec/process"
	"github.com/e2b-dev/infra/packages/envd/internal/utils"

	"connectrpc.com/connect"
	"github.com/creack/pty"
	"github.com/rs/zerolog"
)

const (
	defaultOomScore  = 100
	outputBufferSize = 64
	stdChunkSize     = 2 << 14
	ptyChunkSize     = 2 << 13
)

type ProcessExit struct {
	Error  *string
	Status string
	Exited bool
	Code   int32
}

type Handler struct {
	Config *rpc.ProcessConfig

	logger *zerolog.Logger

	Tag *string
	cmd *exec.Cmd
	tty *os.File

	SignalChan chan syscall.Signal

	outWg *sync.WaitGroup
	stdin io.WriteCloser

	DataEvent *MultiplexedChannel[rpc.ProcessEvent_Data]
	EndEvent  *MultiplexedChannel[rpc.ProcessEvent_End]
}

// This method must be called only after the process has been started
func (p *Handler) Pid() uint32 {
	return uint32(p.cmd.Process.Pid)
}

func New(ctx context.Context, user *user.User, req *rpc.StartRequest, logger *zerolog.Logger, envVars *utils.Map[string, string]) (*Handler, error) {
	cmd := exec.CommandContext(ctx, req.GetProcess().GetCmd(), req.GetProcess().GetArgs()...)

	uid, gid, err := permissions.GetUserIds(user)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	signalChan := make(chan syscall.Signal, 1)
	shutdownChan := make(chan bool)

	go func() {
		signal := <-signalChan
		if signal == syscall.SIGKILL || signal == syscall.SIGTERM {
			// closing the shutdownChan allows us to broadcast the signal to all the read loops below
			close(shutdownChan)
		}
	}()

	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{
		Uid:         uid,
		Gid:         gid,
		Groups:      []uint32{gid},
		NoSetGroups: true,
	}

	resolvedPath, err := permissions.ExpandAndResolve(req.GetProcess().GetCwd(), user)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	cmd.Dir = resolvedPath

	var formattedVars []string

	// Take only 'PATH' variable from the current environment
	// The 'PATH' should ideally be set in the environment
	formattedVars = append(formattedVars, "PATH="+os.Getenv("PATH"))
	formattedVars = append(formattedVars, "HOME="+user.HomeDir)
	formattedVars = append(formattedVars, "USER="+user.Username)
	formattedVars = append(formattedVars, "LOGNAME="+user.Username)

	// Add the environment variables from the global environment
	if envVars != nil {
		envVars.Range(func(key string, value string) bool {
			formattedVars = append(formattedVars, key+"="+value)
			return true
		})
	}

	// Only the last values of the env vars are used - this allows for overwriting defaults
	for key, value := range req.GetProcess().GetEnvs() {
		formattedVars = append(formattedVars, key+"="+value)
	}

	cmd.Env = formattedVars

	outMultiplex := NewMultiplexedChannel[rpc.ProcessEvent_Data](outputBufferSize)
	var outWg sync.WaitGroup

	if req.GetPty() != nil {
		// The pty should ideally start only in the Start method, but the package does not support that and we would have to code it manually.
		// The output of the pty should correctly be passed though.
		tty, err := pty.StartWithSize(cmd, &pty.Winsize{
			Cols: uint16(req.GetPty().GetSize().Cols),
			Rows: uint16(req.GetPty().GetSize().Rows),
		})
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("error starting pty with command '%s' in dir '%s' with '%d' cols and '%d' rows: %w", cmd, cmd.Dir, req.GetPty().GetSize().Cols, req.GetPty().GetSize().Rows, err))
		}

		outWg.Add(1)

		go func() {
			defer outWg.Done()

			for {
				buf := make([]byte, ptyChunkSize)

				n, readErr := tty.Read(buf)

				if n > 0 {
					outMultiplex.Source <- rpc.ProcessEvent_Data{
						Data: &rpc.ProcessEvent_DataEvent{
							Output: &rpc.ProcessEvent_DataEvent_Pty{
								Pty: buf[:n],
							},
						},
					}
				}

				if errors.Is(readErr, io.EOF) {
					break
				}

				if readErr != nil {
					fmt.Fprintf(os.Stderr, "error reading from pty: %s\n", readErr)

					break
				}
			}
		}()

		return &Handler{
			Config:    req.GetProcess(),
			cmd:       cmd,
			tty:       tty,
			Tag:       req.Tag,
			DataEvent: outMultiplex,
			outWg:     &outWg,
			EndEvent:  NewMultiplexedChannel[rpc.ProcessEvent_End](0),
			logger:    logger,
		}, nil
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error creating stdin pipe for command '%s': %w", cmd, err))
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error creating stdout pipe for command '%s': %w", cmd, err))
	}

	outWg.Add(1)
	go func() {
		stdoutLogs := make(chan []byte, outputBufferSize)
		defer close(stdoutLogs)

		readOver := make(chan bool, 1)
		stdoutLogger := logger.With().Str("event_type", "stdout").Logger()

		go logs.LogBufferedDataEvents(stdoutLogs, &stdoutLogger, "data")

		go func() {
			select {
			case <-readOver:
				outWg.Done()
				break
			case _, ok := <-shutdownChan:
				if !ok {
					outWg.Done()
				}
				break
			}
		}()

		for {
			buf := make([]byte, stdChunkSize)

			n, readErr := stdout.Read(buf)

			if n > 0 {
				outMultiplex.Source <- rpc.ProcessEvent_Data{
					Data: &rpc.ProcessEvent_DataEvent{
						Output: &rpc.ProcessEvent_DataEvent_Stdout{
							Stdout: buf[:n],
						},
					},
				}

				stdoutLogs <- buf[:n]
			}

			if errors.Is(readErr, io.EOF) {
				break
			}

			if readErr != nil {
				fmt.Fprintf(os.Stderr, "error reading from stdout: %s\n", readErr)

				break
			}
		}

		readOver <- true
	}()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("error creating stderr pipe for command '%s': %w", cmd, err))
	}

	outWg.Add(1)
	go func() {
		stderrLogs := make(chan []byte, outputBufferSize)
		defer close(stderrLogs)

		readOver := make(chan bool, 1)

		stderrLogger := logger.With().Str("event_type", "stderr").Logger()

		go logs.LogBufferedDataEvents(stderrLogs, &stderrLogger, "data")

		go func() {
			select {
			case <-readOver:
				outWg.Done()
				break
			case _, ok := <-shutdownChan:
				if !ok {
					outWg.Done()
				}
				break
			}
		}()

		for {
			buf := make([]byte, stdChunkSize)

			n, readErr := stderr.Read(buf)

			if n > 0 {
				outMultiplex.Source <- rpc.ProcessEvent_Data{
					Data: &rpc.ProcessEvent_DataEvent{
						Output: &rpc.ProcessEvent_DataEvent_Stderr{
							Stderr: buf[:n],
						},
					},
				}

				stderrLogs <- buf[:n]
			}

			if errors.Is(readErr, io.EOF) {
				break
			}

			if readErr != nil {
				fmt.Fprintf(os.Stderr, "error reading from stderr: %s\n", readErr)

				break
			}
		}

		readOver <- true
	}()

	return &Handler{
		Config:     req.GetProcess(),
		cmd:        cmd,
		stdin:      stdin,
		Tag:        req.Tag,
		DataEvent:  outMultiplex,
		outWg:      &outWg,
		EndEvent:   NewMultiplexedChannel[rpc.ProcessEvent_End](0),
		SignalChan: signalChan,
		logger:     logger,
	}, nil
}

func (p *Handler) SendSignal(signal syscall.Signal) error {
	if p.cmd.Process == nil {
		return fmt.Errorf("process not started")
	}

	return p.cmd.Process.Signal(signal)
}

func (p *Handler) ResizeTty(size *pty.Winsize) error {
	if p.tty == nil {
		return fmt.Errorf("tty not assigned to process")
	}

	return pty.Setsize(p.tty, size)
}

func (p *Handler) WriteStdin(data []byte) error {
	if p.tty != nil {
		return fmt.Errorf("tty assigned to process — input should be written to the pty, not the stdin")
	}

	_, err := p.stdin.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to stdin of process '%d': %w", p.cmd.Process.Pid, err)
	}

	return nil
}

func (p *Handler) WriteTty(data []byte) error {
	if p.tty == nil {
		return fmt.Errorf("tty not assigned to process — input should be written to the stdin, not the tty")
	}

	_, err := p.tty.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to tty of process '%d': %w", p.cmd.Process.Pid, err)
	}

	return nil
}

func (p *Handler) Start() (uint32, error) {
	// Pty is already started in the New method
	if p.tty == nil {
		err := p.cmd.Start()
		if err != nil {
			return 0, fmt.Errorf("error starting process '%s': %w", p.cmd, err)
		}
	}

	adjustErr := adjustOomScore(p.cmd.Process.Pid, defaultOomScore)
	if adjustErr != nil {
		fmt.Fprintf(os.Stderr, "error adjusting oom score for process '%s': %s\n", p.cmd, adjustErr)
	}

	p.logger.
		Info().
		Str("event_type", "process_start").
		Int("pid", p.cmd.Process.Pid).
		Str("command", p.cmd.String()).
		Send()

	return uint32(p.cmd.Process.Pid), nil
}

func (p *Handler) Wait() {
	p.outWg.Wait()

	close(p.DataEvent.Source)

	p.tty.Close()

	err := p.cmd.Wait()

	var errMsg *string

	if err != nil {
		msg := err.Error()
		errMsg = &msg
	}

	endEvent := &rpc.ProcessEvent_EndEvent{
		Error:    errMsg,
		ExitCode: int32(p.cmd.ProcessState.ExitCode()),
		Exited:   p.cmd.ProcessState.Exited(),
		Status:   p.cmd.ProcessState.String(),
	}

	event := rpc.ProcessEvent_End{
		End: endEvent,
	}

	p.EndEvent.Source <- event

	p.logger.
		Info().
		Str("event_type", "process_end").
		Interface("process_result", endEvent).
		Send()
}

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: filesystem/filesystem.proto

package filesystemconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	filesystem "github.com/e2b-dev/infra/packages/envd/internal/services/spec/filesystem"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// FilesystemName is the fully-qualified name of the Filesystem service.
	FilesystemName = "filesystem.Filesystem"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// FilesystemStatProcedure is the fully-qualified name of the Filesystem's Stat RPC.
	FilesystemStatProcedure = "/filesystem.Filesystem/Stat"
	// FilesystemMakeDirProcedure is the fully-qualified name of the Filesystem's MakeDir RPC.
	FilesystemMakeDirProcedure = "/filesystem.Filesystem/MakeDir"
	// FilesystemMoveProcedure is the fully-qualified name of the Filesystem's Move RPC.
	FilesystemMoveProcedure = "/filesystem.Filesystem/Move"
	// FilesystemListDirProcedure is the fully-qualified name of the Filesystem's ListDir RPC.
	FilesystemListDirProcedure = "/filesystem.Filesystem/ListDir"
	// FilesystemRemoveProcedure is the fully-qualified name of the Filesystem's Remove RPC.
	FilesystemRemoveProcedure = "/filesystem.Filesystem/Remove"
	// FilesystemWatchDirProcedure is the fully-qualified name of the Filesystem's WatchDir RPC.
	FilesystemWatchDirProcedure = "/filesystem.Filesystem/WatchDir"
	// FilesystemWatchDirStartProcedure is the fully-qualified name of the Filesystem's WatchDirStart
	// RPC.
	FilesystemWatchDirStartProcedure = "/filesystem.Filesystem/WatchDirStart"
	// FilesystemWatchDirGetProcedure is the fully-qualified name of the Filesystem's WatchDirGet RPC.
	FilesystemWatchDirGetProcedure = "/filesystem.Filesystem/WatchDirGet"
	// FilesystemWatchDirStopProcedure is the fully-qualified name of the Filesystem's WatchDirStop RPC.
	FilesystemWatchDirStopProcedure = "/filesystem.Filesystem/WatchDirStop"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	filesystemServiceDescriptor             = filesystem.File_filesystem_filesystem_proto.Services().ByName("Filesystem")
	filesystemStatMethodDescriptor          = filesystemServiceDescriptor.Methods().ByName("Stat")
	filesystemMakeDirMethodDescriptor       = filesystemServiceDescriptor.Methods().ByName("MakeDir")
	filesystemMoveMethodDescriptor          = filesystemServiceDescriptor.Methods().ByName("Move")
	filesystemListDirMethodDescriptor       = filesystemServiceDescriptor.Methods().ByName("ListDir")
	filesystemRemoveMethodDescriptor        = filesystemServiceDescriptor.Methods().ByName("Remove")
	filesystemWatchDirMethodDescriptor      = filesystemServiceDescriptor.Methods().ByName("WatchDir")
	filesystemWatchDirStartMethodDescriptor = filesystemServiceDescriptor.Methods().ByName("WatchDirStart")
	filesystemWatchDirGetMethodDescriptor   = filesystemServiceDescriptor.Methods().ByName("WatchDirGet")
	filesystemWatchDirStopMethodDescriptor  = filesystemServiceDescriptor.Methods().ByName("WatchDirStop")
)

// FilesystemClient is a client for the filesystem.Filesystem service.
type FilesystemClient interface {
	Stat(context.Context, *connect.Request[filesystem.StatRequest]) (*connect.Response[filesystem.StatResponse], error)
	MakeDir(context.Context, *connect.Request[filesystem.MakeDirRequest]) (*connect.Response[filesystem.MakeDirResponse], error)
	Move(context.Context, *connect.Request[filesystem.MoveRequest]) (*connect.Response[filesystem.MoveResponse], error)
	ListDir(context.Context, *connect.Request[filesystem.ListDirRequest]) (*connect.Response[filesystem.ListDirResponse], error)
	Remove(context.Context, *connect.Request[filesystem.RemoveRequest]) (*connect.Response[filesystem.RemoveResponse], error)
	WatchDir(context.Context, *connect.Request[filesystem.WatchDirRequest]) (*connect.ServerStreamForClient[filesystem.WatchDirResponse], error)
	// Non-streaming versions of WatchDir
	WatchDirStart(context.Context, *connect.Request[filesystem.WatchDirRequest]) (*connect.Response[filesystem.WatchDirStartResponse], error)
	WatchDirGet(context.Context, *connect.Request[filesystem.WatchDirGetRequest]) (*connect.Response[filesystem.WatchDirGetResponse], error)
	WatchDirStop(context.Context, *connect.Request[filesystem.WatchDirStopRequest]) (*connect.Response[filesystem.WatchDirStopResponse], error)
}

// NewFilesystemClient constructs a client for the filesystem.Filesystem service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFilesystemClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) FilesystemClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &filesystemClient{
		stat: connect.NewClient[filesystem.StatRequest, filesystem.StatResponse](
			httpClient,
			baseURL+FilesystemStatProcedure,
			connect.WithSchema(filesystemStatMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		makeDir: connect.NewClient[filesystem.MakeDirRequest, filesystem.MakeDirResponse](
			httpClient,
			baseURL+FilesystemMakeDirProcedure,
			connect.WithSchema(filesystemMakeDirMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		move: connect.NewClient[filesystem.MoveRequest, filesystem.MoveResponse](
			httpClient,
			baseURL+FilesystemMoveProcedure,
			connect.WithSchema(filesystemMoveMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listDir: connect.NewClient[filesystem.ListDirRequest, filesystem.ListDirResponse](
			httpClient,
			baseURL+FilesystemListDirProcedure,
			connect.WithSchema(filesystemListDirMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		remove: connect.NewClient[filesystem.RemoveRequest, filesystem.RemoveResponse](
			httpClient,
			baseURL+FilesystemRemoveProcedure,
			connect.WithSchema(filesystemRemoveMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		watchDir: connect.NewClient[filesystem.WatchDirRequest, filesystem.WatchDirResponse](
			httpClient,
			baseURL+FilesystemWatchDirProcedure,
			connect.WithSchema(filesystemWatchDirMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		watchDirStart: connect.NewClient[filesystem.WatchDirRequest, filesystem.WatchDirStartResponse](
			httpClient,
			baseURL+FilesystemWatchDirStartProcedure,
			connect.WithSchema(filesystemWatchDirStartMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		watchDirGet: connect.NewClient[filesystem.WatchDirGetRequest, filesystem.WatchDirGetResponse](
			httpClient,
			baseURL+FilesystemWatchDirGetProcedure,
			connect.WithSchema(filesystemWatchDirGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		watchDirStop: connect.NewClient[filesystem.WatchDirStopRequest, filesystem.WatchDirStopResponse](
			httpClient,
			baseURL+FilesystemWatchDirStopProcedure,
			connect.WithSchema(filesystemWatchDirStopMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// filesystemClient implements FilesystemClient.
type filesystemClient struct {
	stat          *connect.Client[filesystem.StatRequest, filesystem.StatResponse]
	makeDir       *connect.Client[filesystem.MakeDirRequest, filesystem.MakeDirResponse]
	move          *connect.Client[filesystem.MoveRequest, filesystem.MoveResponse]
	listDir       *connect.Client[filesystem.ListDirRequest, filesystem.ListDirResponse]
	remove        *connect.Client[filesystem.RemoveRequest, filesystem.RemoveResponse]
	watchDir      *connect.Client[filesystem.WatchDirRequest, filesystem.WatchDirResponse]
	watchDirStart *connect.Client[filesystem.WatchDirRequest, filesystem.WatchDirStartResponse]
	watchDirGet   *connect.Client[filesystem.WatchDirGetRequest, filesystem.WatchDirGetResponse]
	watchDirStop  *connect.Client[filesystem.WatchDirStopRequest, filesystem.WatchDirStopResponse]
}

// Stat calls filesystem.Filesystem.Stat.
func (c *filesystemClient) Stat(ctx context.Context, req *connect.Request[filesystem.StatRequest]) (*connect.Response[filesystem.StatResponse], error) {
	return c.stat.CallUnary(ctx, req)
}

// MakeDir calls filesystem.Filesystem.MakeDir.
func (c *filesystemClient) MakeDir(ctx context.Context, req *connect.Request[filesystem.MakeDirRequest]) (*connect.Response[filesystem.MakeDirResponse], error) {
	return c.makeDir.CallUnary(ctx, req)
}

// Move calls filesystem.Filesystem.Move.
func (c *filesystemClient) Move(ctx context.Context, req *connect.Request[filesystem.MoveRequest]) (*connect.Response[filesystem.MoveResponse], error) {
	return c.move.CallUnary(ctx, req)
}

// ListDir calls filesystem.Filesystem.ListDir.
func (c *filesystemClient) ListDir(ctx context.Context, req *connect.Request[filesystem.ListDirRequest]) (*connect.Response[filesystem.ListDirResponse], error) {
	return c.listDir.CallUnary(ctx, req)
}

// Remove calls filesystem.Filesystem.Remove.
func (c *filesystemClient) Remove(ctx context.Context, req *connect.Request[filesystem.RemoveRequest]) (*connect.Response[filesystem.RemoveResponse], error) {
	return c.remove.CallUnary(ctx, req)
}

// WatchDir calls filesystem.Filesystem.WatchDir.
func (c *filesystemClient) WatchDir(ctx context.Context, req *connect.Request[filesystem.WatchDirRequest]) (*connect.ServerStreamForClient[filesystem.WatchDirResponse], error) {
	return c.watchDir.CallServerStream(ctx, req)
}

// WatchDirStart calls filesystem.Filesystem.WatchDirStart.
func (c *filesystemClient) WatchDirStart(ctx context.Context, req *connect.Request[filesystem.WatchDirRequest]) (*connect.Response[filesystem.WatchDirStartResponse], error) {
	return c.watchDirStart.CallUnary(ctx, req)
}

// WatchDirGet calls filesystem.Filesystem.WatchDirGet.
func (c *filesystemClient) WatchDirGet(ctx context.Context, req *connect.Request[filesystem.WatchDirGetRequest]) (*connect.Response[filesystem.WatchDirGetResponse], error) {
	return c.watchDirGet.CallUnary(ctx, req)
}

// WatchDirStop calls filesystem.Filesystem.WatchDirStop.
func (c *filesystemClient) WatchDirStop(ctx context.Context, req *connect.Request[filesystem.WatchDirStopRequest]) (*connect.Response[filesystem.WatchDirStopResponse], error) {
	return c.watchDirStop.CallUnary(ctx, req)
}

// FilesystemHandler is an implementation of the filesystem.Filesystem service.
type FilesystemHandler interface {
	Stat(context.Context, *connect.Request[filesystem.StatRequest]) (*connect.Response[filesystem.StatResponse], error)
	MakeDir(context.Context, *connect.Request[filesystem.MakeDirRequest]) (*connect.Response[filesystem.MakeDirResponse], error)
	Move(context.Context, *connect.Request[filesystem.MoveRequest]) (*connect.Response[filesystem.MoveResponse], error)
	ListDir(context.Context, *connect.Request[filesystem.ListDirRequest]) (*connect.Response[filesystem.ListDirResponse], error)
	Remove(context.Context, *connect.Request[filesystem.RemoveRequest]) (*connect.Response[filesystem.RemoveResponse], error)
	WatchDir(context.Context, *connect.Request[filesystem.WatchDirRequest], *connect.ServerStream[filesystem.WatchDirResponse]) error
	// Non-streaming versions of WatchDir
	WatchDirStart(context.Context, *connect.Request[filesystem.WatchDirRequest]) (*connect.Response[filesystem.WatchDirStartResponse], error)
	WatchDirGet(context.Context, *connect.Request[filesystem.WatchDirGetRequest]) (*connect.Response[filesystem.WatchDirGetResponse], error)
	WatchDirStop(context.Context, *connect.Request[filesystem.WatchDirStopRequest]) (*connect.Response[filesystem.WatchDirStopResponse], error)
}

// NewFilesystemHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFilesystemHandler(svc FilesystemHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	filesystemStatHandler := connect.NewUnaryHandler(
		FilesystemStatProcedure,
		svc.Stat,
		connect.WithSchema(filesystemStatMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	filesystemMakeDirHandler := connect.NewUnaryHandler(
		FilesystemMakeDirProcedure,
		svc.MakeDir,
		connect.WithSchema(filesystemMakeDirMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	filesystemMoveHandler := connect.NewUnaryHandler(
		FilesystemMoveProcedure,
		svc.Move,
		connect.WithSchema(filesystemMoveMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	filesystemListDirHandler := connect.NewUnaryHandler(
		FilesystemListDirProcedure,
		svc.ListDir,
		connect.WithSchema(filesystemListDirMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	filesystemRemoveHandler := connect.NewUnaryHandler(
		FilesystemRemoveProcedure,
		svc.Remove,
		connect.WithSchema(filesystemRemoveMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	filesystemWatchDirHandler := connect.NewServerStreamHandler(
		FilesystemWatchDirProcedure,
		svc.WatchDir,
		connect.WithSchema(filesystemWatchDirMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	filesystemWatchDirStartHandler := connect.NewUnaryHandler(
		FilesystemWatchDirStartProcedure,
		svc.WatchDirStart,
		connect.WithSchema(filesystemWatchDirStartMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	filesystemWatchDirGetHandler := connect.NewUnaryHandler(
		FilesystemWatchDirGetProcedure,
		svc.WatchDirGet,
		connect.WithSchema(filesystemWatchDirGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	filesystemWatchDirStopHandler := connect.NewUnaryHandler(
		FilesystemWatchDirStopProcedure,
		svc.WatchDirStop,
		connect.WithSchema(filesystemWatchDirStopMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/filesystem.Filesystem/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case FilesystemStatProcedure:
			filesystemStatHandler.ServeHTTP(w, r)
		case FilesystemMakeDirProcedure:
			filesystemMakeDirHandler.ServeHTTP(w, r)
		case FilesystemMoveProcedure:
			filesystemMoveHandler.ServeHTTP(w, r)
		case FilesystemListDirProcedure:
			filesystemListDirHandler.ServeHTTP(w, r)
		case FilesystemRemoveProcedure:
			filesystemRemoveHandler.ServeHTTP(w, r)
		case FilesystemWatchDirProcedure:
			filesystemWatchDirHandler.ServeHTTP(w, r)
		case FilesystemWatchDirStartProcedure:
			filesystemWatchDirStartHandler.ServeHTTP(w, r)
		case FilesystemWatchDirGetProcedure:
			filesystemWatchDirGetHandler.ServeHTTP(w, r)
		case FilesystemWatchDirStopProcedure:
			filesystemWatchDirStopHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedFilesystemHandler returns CodeUnimplemented from all methods.
type UnimplementedFilesystemHandler struct{}

func (UnimplementedFilesystemHandler) Stat(context.Context, *connect.Request[filesystem.StatRequest]) (*connect.Response[filesystem.StatResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.Stat is not implemented"))
}

func (UnimplementedFilesystemHandler) MakeDir(context.Context, *connect.Request[filesystem.MakeDirRequest]) (*connect.Response[filesystem.MakeDirResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.MakeDir is not implemented"))
}

func (UnimplementedFilesystemHandler) Move(context.Context, *connect.Request[filesystem.MoveRequest]) (*connect.Response[filesystem.MoveResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.Move is not implemented"))
}

func (UnimplementedFilesystemHandler) ListDir(context.Context, *connect.Request[filesystem.ListDirRequest]) (*connect.Response[filesystem.ListDirResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.ListDir is not implemented"))
}

func (UnimplementedFilesystemHandler) Remove(context.Context, *connect.Request[filesystem.RemoveRequest]) (*connect.Response[filesystem.RemoveResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.Remove is not implemented"))
}

func (UnimplementedFilesystemHandler) WatchDir(context.Context, *connect.Request[filesystem.WatchDirRequest], *connect.ServerStream[filesystem.WatchDirResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.WatchDir is not implemented"))
}

func (UnimplementedFilesystemHandler) WatchDirStart(context.Context, *connect.Request[filesystem.WatchDirRequest]) (*connect.Response[filesystem.WatchDirStartResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.WatchDirStart is not implemented"))
}

func (UnimplementedFilesystemHandler) WatchDirGet(context.Context, *connect.Request[filesystem.WatchDirGetRequest]) (*connect.Response[filesystem.WatchDirGetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.WatchDirGet is not implemented"))
}

func (UnimplementedFilesystemHandler) WatchDirStop(context.Context, *connect.Request[filesystem.WatchDirStopRequest]) (*connect.Response[filesystem.WatchDirStopResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.WatchDirStop is not implemented"))
}

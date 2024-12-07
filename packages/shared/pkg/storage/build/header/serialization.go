package header

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/google/uuid"
)

type Metadata struct {
	Version   int64
	BlockSize int64
	Size      int64
	BuildId   uuid.UUID
}

// Start, Length and SourceStart are in bytes of the data file
// Length will be a multiple of BlockSize
// The list of block mappings will be in order of increasing Start, covering the entire file
type buildMap struct {
	// Offset defines which block of the current layer this mapping starts at
	Offset             uint64
	Length             uint64
	BuildId            uuid.UUID
	BuildStorageOffset uint64
}

func Serialize(metadata *Metadata, mappings []*buildMap, out io.Writer) error {
	err := binary.Write(out, binary.LittleEndian, metadata)
	if err != nil {
		return fmt.Errorf("failed to write metadata: %w", err)
	}

	for _, mapping := range mappings {
		err := binary.Write(out, binary.LittleEndian, mapping)
		if err != nil {
			return fmt.Errorf("failed to write block mapping: %w", err)
		}
	}

	return nil
}

func Deserialize(in io.WriterTo) (*Header, error) {
	var buf bytes.Buffer

	_, err := in.WriteTo(&buf)
	if err != nil {
		return nil, fmt.Errorf("failed to write to buffer: %w", err)
	}

	reader := bytes.NewReader(buf.Bytes())

	var metadata Metadata

	err = binary.Read(reader, binary.LittleEndian, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata: %w", err)
	}

	mappings := make([]*buildMap, 0)

	for {
		var m buildMap
		err := binary.Read(reader, binary.LittleEndian, &m)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("failed to read block mapping: %w", err)
		}

		mappings = append(mappings, &m)
	}

	return NewHeader(&metadata, mappings), nil
}

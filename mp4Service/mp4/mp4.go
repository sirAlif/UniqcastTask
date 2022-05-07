package mp4

import (
	"fmt"
	"io"
	"os"
)

const (
	_FtypBox = "ftyp"
	_MoovBox = "moov"
)

// Box defines a Box structure in mp4 file.
type box struct {
	Type  string
	Size  int64
	Start int64
}

// Reader defines an mp4 structure.
type Reader struct {
	file  io.ReaderAt
	boxes []box
}

func New(path string) (*Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return parse(file)
}

func (r *Reader) GetInitSegment() ([]byte, error) {
	var data []byte
	for _, b := range r.boxes {
		if b.Type == _FtypBox || b.Type == _MoovBox {
			boxData, err := r.readByteAt(b.Size, b.Start)
			if err != nil {
				return nil, err
			}
			data = append(data, boxData...)
		}
		fmt.Println(b.Type, b.Size, b.Start)
	}

	return data, nil
}

// readByteAt reads n byte at offset.
func (r *Reader) readByteAt(n int64, offset int64) ([]byte, error) {
	data := make([]byte, n)
	if _, err := r.file.ReadAt(data, offset); err != nil {
		return nil, err
	}
	return data, nil
}

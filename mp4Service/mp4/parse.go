package mp4

import (
	"encoding/binary"
	"io"
	"os"
)

const (
	_BoxHeaderSize = 8
)

func parse(file *os.File) (*Reader, error) {
	fileSize, err := getSize(file)
	if err != nil {
		return nil, err
	}

	boxes, err := readBoxes(file, fileSize)
	if err != nil {
		return nil, err
	}

	result := &Reader{
		file: file,
		boxes: boxes,
	}

	return result, nil
}

func getSize(file *os.File) (int64, error) {
	info, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}

// readBoxes read boxes for a mp4 file
func readBoxes(file io.ReaderAt, totalSize int64) (boxes []box, err error) {
	for offset := int64(0); offset < totalSize; {
		boxSize, boxType, err := readBoxAt(file, offset)
		if err != nil {
			return nil, err
		}
		b := box{
			Type:  boxType,
			Size:  boxSize,
			Start: offset,
		}
		boxes = append(boxes, b)
		offset += boxSize
	}

	return boxes, nil
}

// readBoxAt reads a box from an offset.
func readBoxAt(file io.ReaderAt, offset int64) (boxSize int64, boxType string, err error) {
	data, err := readHeaderAt(file, offset)
	if err != nil {
		return 0, "", err
	}

	boxSize = int64(binary.BigEndian.Uint32(data[0:4]))
	boxType = string(data[4:8])

	return boxSize, boxType, nil
}

// readHeaderAt reads header at offset.
func readHeaderAt(file io.ReaderAt, offset int64) ([]byte, error) {
	data := make([]byte, _BoxHeaderSize)
	if _, err := file.ReadAt(data, offset); err != nil {
		return nil, err
	}
	return data, nil
}

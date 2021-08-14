// package impl get, insert, delete operations for log.
package log

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	DefaultIndexFileName = ".index.log" // 存放索引的文件名.
	DefaultDataFileName  = ".data.log"  // 存放数据的文件名.
)

// ReadWriter do read and write opreations for log.
type ReadWriter interface {
	// Write a log, caller should maintain sequenceNumber is incremented.
	Write(sequenceNumber uint64, data []byte) error
	// ListBefore all data, that sequenceNumber is less than or equal to given sequenceNumber.
	ListBefore(sequenceNumber uint64) ([][]byte, error)
	// ListAfter all data, that sequenceNumber is larger than or equal to given sequenceNumber.
	ListAfter(sequenceNumber uint64) ([][]byte, error)
	// DropBefore all data, that sequenceNumber is less than or equal to given sequenceNumber.
	DropBefore(sequenceNumber uint64) (int, error)
}

// NewReadWriter create a ReadWriter from given path.
func NewReadWriter(path string) (ReadWriter, error) {
	if err := createDirIfNotExist(path); err != nil {
		return nil, err
	}

	indexFile, err := os.OpenFile(filepath.Join(path, DefaultIndexFileName), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	dataFile, err := os.OpenFile(filepath.Join(path, DefaultDataFileName), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	res := &logReadWriterImpl{
		index: indexFile,
		data:  dataFile,
	}

	return res, nil
}

type logReadWriterImpl struct {
	index *os.File
	data  *os.File
}

// Write a log, caller should maintain sequenceNumber is incremented.
func (*logReadWriterImpl) Write(sequenceNumber uint64, data []byte) error {
	// TODO: Add the index to the index file, and add the data to the Data file.

	return nil
}

// ListBefore all data, that sequenceNumber is less than or equal to given sequenceNumber.
func (*logReadWriterImpl) ListBefore(sequenceNumber uint64) ([][]byte, error) {
	// TODO: Search in the index file, take it out in the data file.

	return nil, nil
}

// ListAfter all data, that sequenceNumber is larger than or equal to given sequenceNumber.
func (*logReadWriterImpl) ListAfter(sequenceNumber uint64) ([][]byte, error) {
	// TODO: Search in the index file, take it out in the data file.

	return nil, nil
}

// DropBefore all data, that sequenceNumber is less than or equal to given sequenceNumber.
func (*logReadWriterImpl) DropBefore(sequenceNumber uint64) (int, error) {
	// TODO: Returns the number of deleted data.

	return 0, nil
}

func createDirIfNotExist(path string) error {
	obj, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return os.Mkdir(path, os.ModePerm)
	}

	if err != nil {
		return err
	}

	// not dir, return error.
	if !obj.IsDir() {
		return fmt.Errorf("%s is not dir", path)
	}

	return nil
}

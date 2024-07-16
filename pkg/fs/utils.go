package fs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Mkdir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func OpenFile(path string) (*os.File, error) {
	dir := filepath.Dir(path)

	err := Mkdir(dir)
	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("file %s already exists", path)
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

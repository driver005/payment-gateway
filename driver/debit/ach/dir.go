package ach

import (
	"fmt"
	"os"
	"path/filepath"
)

// ReadDir will attempt to parse all ACH files in the given directory. Only files which
// parse successfully will be returned.
func ReadDir(dir string) ([]*File, error) {
	readACH := func(path string) (*File, error) {
		fd, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("opening %s failed: %v", path, err)
		}
		defer fd.Close()

		f, err := NewReader(fd).Read()
		if err != nil {
			return nil, fmt.Errorf("reading %s failed: %v", path, err)
		}
		return &f, nil
	}

	readJSON := func(path string) (*File, error) {
		bs, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("opening %s failed: %v", path, err)
		}
		return FileFromJSON(bs)
	}

	infos, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var out []*File
	for i := range infos {
		path := filepath.Join(dir, infos[i].Name())

		f, err1 := readACH(path)
		if f != nil {
			out = append(out, f)
			continue
		}
		f, err2 := readJSON(path)
		if f != nil {
			out = append(out, f)
			continue
		}

		if err1 != nil && err2 != nil {
			return out, fmt.Errorf("%s failed to parse: %v | %v", path, err1, err2)
		}
	}
	return out, nil
}

package backend

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type FSBackend struct {
	data     map[string]string
	filename string
}

func NewFSBackend(filename string) *FSBackend {
	return &FSBackend{
		data:     map[string]string{"gui": "nat"},
		filename: filename,
	}
}

func (be *FSBackend) Load() error {
	file, err := os.OpenFile(
		be.filename,
		os.O_CREATE|os.O_RDWR,
		os.ModePerm,
	)
	if err != nil {
		return err
	}

	fileReader := bufio.NewReader(file)

	var line []byte

	for {
		line, _, err = fileReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		lineString := strings.TrimPrefix(string(line), "\n")
		entry := strings.Split(lineString, "\t")
		if len(entry) < 2 {
			return errors.New("malformed kv entry")
		}

		be.data[entry[0]] = entry[1]
	}

	return nil
}

func (be *FSBackend) Set(key, value string) error {
	be.data[key] = value
	return be.dump()
}

func (be FSBackend) Get(key string) (string, error) {
	value, ok := be.data[key]

	if !ok {
		return "", nil
	}

	return value, nil
}

func (be *FSBackend) Delete(key string) error {
	if _, ok := be.data[key]; ok {
		delete(be.data, key)
		return be.dump()
	}
	return errors.New("Can't delete key : '" + key + "' -> not in kv")
}

func (be FSBackend) dump() error {
	kvFile, err := os.OpenFile(
		be.filename,
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.ModePerm,
	)

	if err != nil {
		return err
	}
	defer kvFile.Close()

	for k, v := range be.data {
		kvFile.WriteString(k + "\t" + v + "\n")
	}

	return nil
}

package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type KV struct {
	data         map[string]string
	dumpFilename string
}

func NewKv(dumpFilename string) *KV {
	return &KV{
		data:         map[string]string{"gui": "nat"},
		dumpFilename: dumpFilename,
	}
}

func (kv *KV) Load() error {
	file, err := os.OpenFile(
		kv.dumpFilename,
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

		kv.data[entry[0]] = entry[1]
	}

	return nil
}

func (kv *KV) Set(key, value string) {
	kv.data[key] = value
}

func (kv KV) Get(key string) string {
	value, ok := kv.data[key]

	if !ok {
		return ""
	}

	return value
}

func (kv *KV) Delete(key string) error {
	if _, ok := kv.data[key]; ok {
		delete(kv.data, key)
		return nil
	}
	return errors.New("Can't delete key : '" + key + "' -> not in kv")
}

func (kv KV) Dump() error {
	kvFile, err := os.OpenFile(
		kv.dumpFilename,
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.ModePerm,
	)

	if err != nil {
		return err
	}
	defer kvFile.Close()

	for k, v := range kv.data {
		kvFile.WriteString(k + "\t" + v + "\n")
	}

	return nil
}

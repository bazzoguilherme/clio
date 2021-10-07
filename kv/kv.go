package kv

import (
	"log"
)

type KV struct {
	backend Backend
}

func NewKv(backend Backend) (*KV, error) {
	err := backend.Load()
	if err != nil {
		return nil, err
	}

	return &KV{
		backend: backend,
	}, nil
}

func (kv *KV) Set(key, value string) {
	err := kv.backend.Set(key, value)
	if err != nil {
		log.Println("Error - Set Method: %s", err)
	}
}

func (kv KV) Get(key string) string {
	value, err := kv.backend.Get(key)

	if err != nil {
		log.Println("Error - Get Method: %s", err)
		return ""
	}

	return value
}

func (kv KV) Delete(key string) {
	err := kv.backend.Delete(key)

	if err != nil {
		log.Println("Error - Delete Method: %s", err)
	}
}

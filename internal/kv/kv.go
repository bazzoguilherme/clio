package kv

import (
	"errors"
	"fmt"
	"log"
)

type KV struct {
	backend Backend
}

func NewKv(backend Backend) (*KV, error) {
	err := backend.Load()
	if err != nil {
		return nil, fmt.Errorf("kv construction failed: %w", ErrBanckendLoadFailed)
	}

	return &KV{
		backend: backend,
	}, nil
}

func (kv *KV) Set(key, value string) {
	err := kv.backend.Set(key, value)
	if err != nil {
		log.Printf("Error - Set Method: %s", err.Error())
	}
}

func (kv KV) Get(key string) string {
	value, err := kv.backend.Get(key)

	if err != nil {
		switch {
		case errors.Is(err, ErrBanckendGetKeyDoesntExist):
			log.Printf("Failed Get: key doesn't exist: %s", err.Error())
		case errors.Is(err, ErrBanckendGetFailed):
			log.Printf("Failed Get: Error in Get: %s", err.Error())
		default:
			log.Printf("Error - Get Method: %s", err.Error())
		}

		return ""
	}

	return value
}

func (kv KV) Delete(key string) {
	err := kv.backend.Delete(key)

	if err != nil {
		log.Printf("Error - Delete Method: %s", err.Error())
	}
}

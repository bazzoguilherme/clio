package kv

import "errors"

var (
	ErrBanckendLoadFailed        = errors.New("backend load failed")
	ErrBanckendSetFailed         = errors.New("backend set failed")
	ErrBanckendGetKeyDoesntExist = errors.New("backend get failed - key doen't exist")
	ErrBanckendGetFailed         = errors.New("backend get failed")
)

type Backend interface {
	Load() error
	Set(string, string) error
	Get(string) (string, error)
	Delete(string) error
}

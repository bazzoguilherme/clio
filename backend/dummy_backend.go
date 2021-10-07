package backend

import "log"

type DummyBackend struct{}

func (be DummyBackend) Load() error {
	log.Println("Dummy Backend = load")
	return nil
}

func (be DummyBackend) Set(string, string) error {
	log.Println("Dummy Backend = set")
	return nil
}

func (be DummyBackend) Get(string) (string, error) {
	log.Println("Dummy Backend = get")
	return "dummy", nil
}

func (be DummyBackend) Delete(string) error {
	log.Println("Dummy Backend = delete")
	return nil
}

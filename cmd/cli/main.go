package main

import (
	"errors"
	"log"
	"os"

	be "github.com/bazzoguilherme/clio/internal/backend"
	keyvalue "github.com/bazzoguilherme/clio/internal/kv"
)

var filename = "./kv.db"

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Fatal("not enough arguments")
	}

	backendType := args[1]
	var backend keyvalue.Backend
	switch backendType {
	case "dummy":
		backend = be.DummyBackend{}
	case "fs":
		backend = be.NewFSBackend(filename)
	case "http":
		backend = be.NewHttpBackend()
	default:
		log.Fatal("Unsupported Backend")
	}

	command := args[2]
	commandArgs := args[3:]

	kv, err := keyvalue.NewKv(backend)
	if err != nil {
		if errors.Is(err, keyvalue.ErrBanckendLoadFailed) {
			log.Fatal("BACKEND LOAD FAILED")
		}
		log.Fatal(err)
	}

	switch command {
	case "set":
		if len(commandArgs) < 2 {
			log.Fatal("not enough arguments for 'set' command")
		}

		key, value := commandArgs[0], commandArgs[1]

		kv.Set(key, value)

	case "get":
		if len(commandArgs) < 1 {
			log.Fatal("not enough arguments for 'get' command")
		}

		key := commandArgs[0]

		value := kv.Get(key)

		log.Println(value)
	case "delete":
		if len(commandArgs) < 1 {
			log.Fatal("not enough arguments for 'delete' command")
		}

		key := commandArgs[0]

		kv.Delete(key)

	default:
		log.Fatal("unsuported command")
	}

}

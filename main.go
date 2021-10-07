package main

import (
	"log"
	"os"

	be "github.com/bazzoguilherme/clio/backend"
	kv "github.com/bazzoguilherme/clio/kv"
)

var filename = "./kv.db"

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("not enough arguments")
	}

	command := args[1]
	commandArgs := args[2:]

	fs_backend := be.NewFSBackend(filename)

	kv, err := kv.NewKv(fs_backend)
	if err != nil {
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

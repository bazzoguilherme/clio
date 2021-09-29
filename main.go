package main

import (
	"log"
	"os"
)

var filename = "./kv.db"

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("not enough arguments")
	}

	command := args[1]
	commandArgs := args[2:]

	kv := NewKv(filename)
	err := kv.Load()
	if err != nil {
		log.Fatal("unable to load kv file")
	}

	switch command {
	case "set":
		if len(commandArgs) < 2 {
			log.Fatal("not enough arguments for 'set' command")
		}

		key, value := commandArgs[0], commandArgs[1]

		kv.Set(key, value)

		err := kv.Dump()
		if err != nil {
			log.Fatal(err)
		}

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

		var err error
		err = kv.Delete(key)
		if err != nil {
			log.Fatal(err)
		}

		err = kv.Dump()
		if err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatal("unsuported command")
	}

}

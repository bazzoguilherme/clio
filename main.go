package main

import (
	"log"
	"os"
)

var data = map[string]string{"gui": "nat"}

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("not enough arguments")
	}

	command := args[1]
	commandArgs := args[2:]

	switch command {
	case "set":
		if len(commandArgs) < 2 {
			log.Fatal("not enough arguments for 'set' command")
		}

		key, value := commandArgs[0], commandArgs[1]

		Set(key, value)

		kvFile, err := os.OpenFile(
			"./kv.db",
			os.O_APPEND|os.O_CREATE|os.O_RDWR,
			os.ModePerm,
		)

		if err != nil {
			log.Fatal(err.Error())
		}
		defer kvFile.Close()

		for k, v := range data {
			kvFile.WriteString(k + "\t" + v + "\n")
		}

	case "get":
		if len(commandArgs) < 1 {
			log.Fatal("not enough arguments for 'get' command")
		}

		key := commandArgs[0]

		value := Get(key)

		log.Println(value)
	default:
		log.Fatal("unsuported command")
	}

}

func Set(key, value string) {
	data[key] = value
}

func Get(key string) string {
	value, ok := data[key]

	if !ok {
		return ""
	}

	return value
}

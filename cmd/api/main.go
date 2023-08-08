package main

import (
	"github.com/ravilock/naive-banking-system/api"
	"log"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Start()
}

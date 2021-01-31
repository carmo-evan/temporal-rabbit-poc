package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

// This package emulates a frontend application sending a request to a server

func main() {
	client := http.Client{}
	res, err := client.Post("http://localhost:1914", "application/json", strings.NewReader(`{"picture": "foo.jpg"}`))
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.ReadFrom(res.Body)
}

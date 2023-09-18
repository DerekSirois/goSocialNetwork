package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Welcome to the Go social network")

	app, err := New()
	if err != nil {
		log.Fatal(err)
		return
	}
	app.Routes()
	app.Run()
}

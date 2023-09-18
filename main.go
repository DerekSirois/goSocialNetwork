package main

import (
	"fmt"
	"goSocialNetwork/app"
	"log"
)

func main() {
	fmt.Println("Welcome to the Go social network")

	a, err := app.New()
	handleError(err)

	err = a.Migrate()
	handleError(err)

	a.Routes()
	a.Run()
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

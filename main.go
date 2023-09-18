package main

import (
	"fmt"
	"goSocialNetwork/app"
	"log"
)

func main() {
	fmt.Println("Welcome to the Go social network")

	app, err := app.New()
	handleError(err)

	err = app.Migrate()
	handleError(err)

	app.Routes()
	app.Run()
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

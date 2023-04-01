package main

import "github.com/lara-store/damo/cmd/httpServer/application"

func main() {
	err := application.New().Listen()

	if err != nil {
		panic(err)
	}
}

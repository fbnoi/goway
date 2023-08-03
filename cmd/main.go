package main

import (
	"flynoob/goway/app/service"
	"flynoob/goway/internal"
	"log"
)

func main() {
	serve := internal.NewServer()
	serve.SetBeforeUpgradeHandler(service.Authenticate)
	serve.SetAfterUpgradeHandler(func(c *internal.Client) {
		// service.
	})
	if err := serve.Run(); err != nil {
		log.Fatal(err)
	}
}

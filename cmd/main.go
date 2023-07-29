package main

import (
	"flynoob/goway"
	"log"
	"net/http"
)

func main() {
	serve := goway.NewServer()
	serve.SetBeforeUpgradeHandler(func(w http.ResponseWriter, r *http.Request) bool {
		log.Println("New Connection come in")
		return true
	})
	serve.SetAfterUpgradeHandler(func(c *goway.Client) {
		log.Println("Connection Established")
	})
	serve.SetPingHandler(func(c *goway.Client, b []byte) {
		log.Printf("Receive ping: %s \n", b)
	})
	serve.SetPongHandler(func(c *goway.Client, b []byte) {
		log.Printf("Receive pong: %s \n", b)
	})
	serve.SetCloseHandler(func(c *goway.Client) {
		log.Println("Connection closed")
	})
	serve.Run()
}

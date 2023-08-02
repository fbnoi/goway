package main

import (
	"flynoob/goway"
	"log"
	"net/http"

	"flynoob/goway/internal"
)

func main() {
	serve := goway.NewServer()
	serve.SetBeforeUpgradeHandler(func(w http.ResponseWriter, r *http.Request) bool {
		log.Println("New Connection come in")
		return true
	})
	serve.SetAfterUpgradeHandler(func(c *goway.Client) {
		log.Println("Connection Established")
		internal.MonitorHealth(c)
	})
	serve.SetByteMessageHandler(func(c *goway.Client, bs []byte) {
		go func() {
			if a, err := internal.GetFrameFromBytes(bs); err != nil {
				log.Println(err)
			} else {
				c.Publish(frame)
				internal.PutFrame(frame)
			}
		}()
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
	if err := serve.Run(); err != nil {
		log.Fatal(err)
	}
}

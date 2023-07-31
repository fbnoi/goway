package main

import (
	"flynoob/goway"
	"log"
	"net/http"

	"flynoob/goway/internal"
	pb "flynoob/goway/protobuf"
)

func main() {
	serve := goway.NewServer()
	serve.SetBeforeUpgradeHandler(func(w http.ResponseWriter, r *http.Request) bool {
		log.Println("New Connection come in")
		return true
	})
	serve.SetAfterUpgradeHandler(func(c *goway.Client) {
		log.Println("Connection Established")
		c.Subscribe(pb.FrameType_HEARTBEAT, func(f *pb.Frame) {
			if err := internal.OnPing(c, f); err != nil {
				log.Printf("Ping error: %s\n", err)
			}
		})
		internal.CheckHealthy(c)
	})
	serve.SetByteMessageHandler(func(c *goway.Client, bs []byte) {
		go func() {
			if frame, err := internal.GetFrameFromBytes(bs); err != nil {
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

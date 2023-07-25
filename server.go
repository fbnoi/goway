package goway

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader websocket.Upgrader

	beforeUpgradeHandler func(w http.ResponseWriter, r *http.Request) bool
	afterUpgradeHandler  func(*Client)
	pingHandler          func(*Client, []byte)
	pongHandler          func(*Client, []byte)
	closeHandler         func(*Client)
	messageHandler       func(*Client, int, []byte)
}

func (s *Server) Listen(addr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !s.beforeUpgradeHandler(w, r) {
			return
		}
		c, err := s.upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		client := &Client{conn: c}
		s.afterUpgradeHandler(client)
		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			switch mt {
			case websocket.PingMessage:
				s.pingHandler(client, message)
			case websocket.PongMessage:
				s.pongHandler(client, message)
			case websocket.TextMessage, websocket.BinaryMessage:
				s.messageHandler(client, mt, message)
			case websocket.CloseMessage:
				s.closeHandler(client)
				return
			}
		}
	})
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (s *Server) SetAfterUpgradeHandler(handler func(*Client)) {
	s.afterUpgradeHandler = handler
}

func (s *Server) SetBeforeUpgradeHandler(handler func(w http.ResponseWriter, r *http.Request) bool) {
	s.beforeUpgradeHandler = handler
}

func (s *Server) SetCloseHandler(handler func(*Client)) {
	s.closeHandler = handler
}

func (s *Server) SetMessageHandler(handler func(*Client, int, []byte)) {
	s.messageHandler = handler
}

func (s *Server) SetPingHandler(handler func(*Client, []byte)) {
	s.pingHandler = handler
}

func (s *Server) SetPongHandler(handler func(*Client, []byte)) {
	s.pongHandler = handler
}

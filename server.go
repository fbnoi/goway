package goway

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	addr               string
	handleShakeTimeout int64
	readBufferSize     int
	writeBufferSize    int
	enableCompression  bool
)

func parseFlag() {
	flag.StringVar(&addr, "addr", ":5678", "set addr for server")
	flag.Int64Var(&handleShakeTimeout, "timeout", 3000, "set connection timeout")
	flag.IntVar(&readBufferSize, "read_buffer_size", 0, "set read buffer size")
	flag.IntVar(&writeBufferSize, "write_buffer_size", 0, "set write buffer size")
	flag.BoolVar(&enableCompression, "enable_compression", false, "enable message compression")
}

func NewServer() *Server {
	parseFlag()
	return &Server{upgrader: websocket.Upgrader{
		HandshakeTimeout:  time.Millisecond * time.Duration(handleShakeTimeout),
		ReadBufferSize:    readBufferSize,
		WriteBufferSize:   writeBufferSize,
		EnableCompression: enableCompression,
		CheckOrigin: func(*http.Request) bool {
			return true
		},
	}, addr: addr}
}

type Server struct {
	addr     string
	upgrader websocket.Upgrader

	beforeUpgrade func(w http.ResponseWriter, r *http.Request) bool
	afterUpgrade  func(*Client)
	handlePing    func(*Client, []byte)
	handlePong    func(*Client, []byte)
	handleClose   func(*Client)
	handleMessage func(*Client, int, []byte)
}

func (s *Server) Run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !s.beforeUpgrade(w, r) {
			return
		}
		c, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		client := &Client{conn: c}
		s.afterUpgrade(client)
		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			switch mt {
			case websocket.PingMessage:
				s.handlePing(client, message)
			case websocket.PongMessage:
				s.handlePong(client, message)
			case websocket.TextMessage, websocket.BinaryMessage:
				s.handleMessage(client, mt, message)
			case websocket.CloseMessage:
				s.handleClose(client)
				return
			}
		}
	})
	return http.ListenAndServe(s.addr, nil)
}

func (s *Server) Listen(addr string) error {
	s.addr = addr
	return s.Run()
}

func (s *Server) SetAfterUpgradeHandler(handler func(*Client)) {
	s.afterUpgrade = handler
}

func (s *Server) SetBeforeUpgradeHandler(handler func(w http.ResponseWriter, r *http.Request) bool) {
	s.beforeUpgrade = handler
}

func (s *Server) SetCloseHandler(handler func(*Client)) {
	s.handleClose = handler
}

func (s *Server) SetMessageHandler(handler func(*Client, int, []byte)) {
	s.handleMessage = handler
}

func (s *Server) SetPingHandler(handler func(*Client, []byte)) {
	s.handlePing = handler
}

func (s *Server) SetPongHandler(handler func(*Client, []byte)) {
	s.handlePong = handler
}

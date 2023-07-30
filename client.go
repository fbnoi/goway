package goway

import (
	pb "flynoob/goway/protobuf"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type Color int
type SocketStatus int

const (
	Green  = Color(1)
	Blue   = Color(2)
	Yellow = Color(3)
	Red    = Color(4)
)

const (
	Connected    = SocketStatus(1)
	DisConnected = SocketStatus(2)
)

func NewClient(serve *Server, conn *websocket.Conn, uid string) *Client {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": serve.endpoint,
		"aud": uid,
		"nbf": time.Now().UnixMilli(),
		"iat": time.Now().UnixMilli(),
	})
	return &Client{
		conn:      conn,
		Color:     Green,
		bus:       NewBus(),
		serve:     serve,
		status:    Connected,
		authToken: token,
	}
}

type Client struct {
	conn    *websocket.Conn
	session []*KV[any]
	status  SocketStatus
	serve   *Server

	LastPingAt time.Time
	Color      Color
	bus        Bus

	authToken *jwt.Token

	sync.RWMutex
}

func (c *Client) Publish(f *pb.Frame) {
	c.bus.Publish(f)
	c.bus.WaitAsync()
}

func (c *Client) Subscribe(typ pb.FrameType, handleFunc func(f *pb.Frame)) {
	c.bus.SubscribeAsync(typ, handleFunc, false)
}

func (c *Client) SubscribeOnce(typ pb.FrameType, handleFunc func(f *pb.Frame)) {
	c.bus.SubscribeOnceAsync(typ, handleFunc)
}

func (c *Client) Send(frame *pb.Frame) error {
	if bs, err := proto.Marshal(frame); err != nil {
		return err
	} else {
		return c.doSend(2, bs)
	}
}

func (c *Client) Close() error {
	c.Lock()
	defer c.Unlock()
	if c.status == DisConnected {
		return nil
	}
	c.status = DisConnected
	c.conn.WriteControl(websocket.CloseMessage, nil, time.Now().Add(time.Second))
	return c.conn.Close()
}

func (c *Client) Status() SocketStatus {
	c.RLock()
	defer c.RUnlock()
	return c.status
}

func (c *Client) Set(name string, val any) {
	c.Lock()
	defer c.Unlock()
	if kv, ok := c.Get(name); ok {
		kv.Value = val
		return
	}
	c.session = append(c.session, &KV[any]{Key: name, Value: val})
}

func (c *Client) Get(name string) (*KV[any], bool) {
	c.RLock()
	defer c.RUnlock()
	for _, kv := range c.session {
		if kv.Key == name {
			return kv, true
		}
	}
	return nil, false
}

func (c *Client) doSend(mt int, message []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.status != Connected {
		return errors.New("Websocket is disconnected")
	}
	return c.conn.WriteMessage(mt, message)
}

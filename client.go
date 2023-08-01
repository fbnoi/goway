package goway

import (
	pb "flynoob/goway/protobuf"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
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
	return &Client{
		conn:      conn,
		Color:     Green,
		bus:       NewBus(),
		serve:     serve,
		status:    Connected,
		authToken: serve.GenClientToken(uid),
	}
}

type Client struct {
	conn    *websocket.Conn
	session map[string]any
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

func (c *Client) Send(message proto.Message) error {
	m, err := anypb.New(message)
	if err != nil {
		return err
	}
	bs, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	return c.doSend(2, bs)
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
	c.session[name] = val
}

func (c *Client) Get(name string) (val any, ok bool) {
	c.RLock()
	defer c.RUnlock()
	val, ok = c.session[name]
	return
}

func (c *Client) Delete(name string) {
	c.Lock()
	defer c.Unlock()
	delete(c.session, name)
}

func (c *Client) doSend(mt int, message []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.status != Connected {
		return errors.New("Websocket is disconnected")
	}
	return c.conn.WriteMessage(mt, message)
}

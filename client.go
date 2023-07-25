package goway

import (
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn    *websocket.Conn
	session []*KV[any]
}

func (c *Client) Send(mt int, message []byte) error {
	return c.conn.WriteMessage(mt, message)
}

func (c *Client) Ping(message []byte) error {
	return c.conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Second))
}

func (c *Client) Close() error {
	c.conn.WriteControl(websocket.CloseMessage, nil, time.Now().Add(time.Second))

	return c.conn.Close()
}

func (c *Client) CloseWithMessage(message []byte) error {
	c.conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))

	return c.conn.Close()
}

func (c *Client) Set(name string, val any) {
	if kv, ok := c.Get(name); ok {
		kv.Value = val

		return
	}
	c.session = append(c.session, &KV[any]{Key: name, Value: val})
}

func (c *Client) Get(name string) (*KV[any], bool) {
	for _, kv := range c.session {
		if kv.Key == name {
			return kv, true
		}
	}

	return nil, false
}

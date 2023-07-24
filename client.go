package goway

import "github.com/gorilla/websocket"

func NewClient(conn *websocket.Conn) *Client {
	return &Client{conn: conn}
}

type Client struct {
	conn    *websocket.Conn
	session []*KV[any]
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

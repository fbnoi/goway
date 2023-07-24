package goway

import (
	"github.com/pkg/errors"
)

type Register struct {
	Gateways []*KV[*Client]
	Workers  []*KV[*Client]
}

func (r *Register) register(id string, client *Client) error {
	for _, kv := range r.Gateways {
		if kv.Key == id {
			return errors.Errorf("Client named %s already exist", id)
		}
	}
	r.Gateways = append(r.Gateways, &KV[*Client]{Key: id, Value: client})

	return nil
}

func (r *Register) broadcast() error {

}

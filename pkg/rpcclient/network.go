package rpcclient

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/olympus-protocol/ogen/api/proto"
)

func (c *Client) GetNetworkInfo() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.network.GetNetworkInfo(ctx, &proto.Empty{})
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Client) GetPeersInfo() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.network.GetPeersInfo(ctx, &proto.Empty{})
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Client) AddPeer(args []string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if len(args) < 1 {
		return "", errors.New("Usage: addpeer <addr>")
	}
	req := &proto.IP{Host: args[0]}
	res, err := c.network.AddPeer(ctx, req)
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

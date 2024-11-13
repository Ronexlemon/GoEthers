package config

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	RPC_URL string
}

func (r *Client) NewClient() *ethclient.Client {
	client, err := ethclient.Dial(r.RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return client
}

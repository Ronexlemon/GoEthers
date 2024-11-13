package services

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/text/number"
)

type RPC  struct{
	RPC_URL string
}


func(rpc *RPC)GetBalance(ctx context,client ethclient,account string)(number){
	balance := client.BalanceAt(ctx,common.HexToAddress(account),nil)
	return balance

} 
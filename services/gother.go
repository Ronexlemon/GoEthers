package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


func GetBalance(ctx context.Context, client *ethclient.Client, account string) (*big.Int, error) {
	
	balance, err := client.BalanceAt(ctx, common.HexToAddress(account), nil)
	if err != nil {
		return nil, err 
	}
	return balance, nil
}

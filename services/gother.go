package services

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


func GetBalance(ctx context.Context, client *ethclient.Client, account string) (*big.Int, error) {
	
	balance, err := client.BalanceAt(ctx, common.HexToAddress(account), nil)
	if err != nil {
		return nil, err 
	}
	fmt.Println("Balance user",balance)
	return balance, nil
}

//get blockNumber
func GetBlockNumber(ctx context.Context, client *ethclient.Client) (*uint64, error){
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, err
		}
		return &blockNumber, nil
}
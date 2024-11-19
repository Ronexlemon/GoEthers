package services

import (
	"context"
	"fmt"

	"math/big"

	"github.com/RonexLemon/Goether/types"
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

func GetLatestBlockDetails(ctx context.Context, client *ethclient.Client) (types.Block, error) {
	var blockcust types.Block

	// Get the latest block (passing nil to get the latest block)
	blockNumber,err := GetBlockNumber(context.Background(),client)
	//recheck
	if err != nil {
		return blockcust, err}
		blockNumberBigInt := new(big.Int).SetUint64(*blockNumber)
block, err := client.BlockByNumber(context.Background(), blockNumberBigInt)
fmt.Println(block.Transactions())
if err != nil {
  return types.Block{},err
}
if block == nil {
	return blockcust, fmt.Errorf("failed to retrieve the block")
}

	
	blockcust = types.Block{
		Hash:        block.Hash(),
		Number:      blockcust.Number,
		
	}

	// Return the pointer to blockcust
	return blockcust, nil
}
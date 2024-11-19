package services

import (
	"context"
	"fmt"

	"log"
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
	fmt.Println("Balance user", balance)
	return balance, nil
}

// get blockNumber
func GetBlockNumber(ctx context.Context, client *ethclient.Client) (*uint64, error) {
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	return &blockNumber, nil
}

func GetLatestBlockDetails(ctx context.Context, client *ethclient.Client) (types.Block, error) {
	var blockcust types.Block
	blocknum, err := GetBlockNumber(ctx, client)
	if err != nil {
		return blockcust, err
	}
	fmt.Println("Current Block", blocknum)

	// Get the latest block (passing nil to get the latest block)
	blockNumber := big.NewInt(31080000)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	fmt.Println("Block details", block)
	if err != nil {
		return blockcust, err

	}

	fmt.Println("Block Number:", block.Number().Uint64())             // 5671744
	fmt.Println("Block Timestamp:", block.Time())                     // 1527211625
	fmt.Println("Block Difficulty:", block.Difficulty().Uint64())     // 3217000136609065
	fmt.Println("Block Hash:", block.Hash().Hex())                    // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println("Number of Transactions:", len(block.Transactions())) // 144// 144

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	blockcust = types.Block{
		Number: block.Number().Uint64(),
		Hash:   block.Hash(),
		//ParentHash:   block.ParentHash(),
		Time:       block.Time(),
		Difficulty: block.Difficulty().Uint64(),
		// Transactions: block.Transactions(),
	}

	fmt.Println("Count", count)
	return blockcust, nil
}

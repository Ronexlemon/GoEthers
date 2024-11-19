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
	blockNumber := big.NewInt(int64(*blocknum))
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	
	if err != nil {
		return blockcust, err

	}

	

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	var transactions []types.Transaction

// Loop through each transaction in the block
for _, tx := range block.Transactions() {
    // Add each transaction to the transactions slice
    transactions = append(transactions, types.Transaction{
		Hash: tx.Hash(),
		Nonce  : tx. Nonce(),
		GasPrice: tx.GasPrice(),
		//GasLimit: tx.GasLimit(),
		To: tx.To(),
		Value: tx.Value(),
		Input: tx.Data(),

		
	})
}

// Construct the blockcust object
blockcust = types.Block{
    Number:      block.Number().Uint64(),
    Hash:        block.Hash(),
    Time:        block.Time(),
    Difficulty:  block.Difficulty().Uint64(),
    Transactions: transactions, // Assign the list of transactions
}

	fmt.Println("Count", count)
	return blockcust, nil
}

func GetTransactionByHash(ctx context.Context,client *ethclient.Client, hash  string)(types.Transaction,bool,error){
	var transaction types.Transaction
	hashit := common.HexToHash(hash)
	tx,isPending, err := client.TransactionByHash(ctx,hashit);
	if err !=nil{
		return transaction,isPending,err
	}
	transaction = types.Transaction{
		Hash: tx.Hash(),
		//BlockNumber: tx.BlockNumber(),
		//BlockHash: tx.BlockHash(),
		//From: tx.From(),
		//To: tx.To(),
		//Value: tx.Value(),
		//Gas: tx.Gas(),
		//GasPrice: tx.GasPrice(),
		//Nonce: tx.Nonce(),
		}
		return transaction,isPending,nil
}
package services

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IsContractOrAddress(ctx context.Context, client *ethclient.Client, address string) (bool, string, error) {
	// Check if the address is a contract
	address_ := common.HexToAddress(address)
	bytcode, err := client.CodeAt(ctx, address_, nil)
	if err != nil {
		return false, "", err
	}
	if len(bytcode) > 0 {
		return true, "IsContract", nil
	} else {
		return true, "IsAddress", nil
	}

}

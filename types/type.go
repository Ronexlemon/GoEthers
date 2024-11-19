package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)
type Block struct {
    Number     uint64        // Block number
    Time       uint64  // Block hash
    Difficulty uint64
	Hash     common.Hash   // Parent block hash
    Timestamp  uint64         // Block timestamp (unix timestamp)
    Transactions []Transaction // Transactions in the block
}
type Transaction struct {
    Hash       common.Hash // Transaction hash
    Nonce      uint64      // Transaction nonce
    BlockHash  common.Hash // Hash of the block containing the transaction
   BlockNumber *big.Int   // Block number
   From       common.Address // Sender address
    To         *common.Address // Receiver address, nil for contract creation
    Value      *big.Int    // Value in wei
    Gas        uint64      // Gas limit
    GasPrice   *big.Int    // Gas price in wei
    Input      []byte      // Input data for smart contracts
    //V          *big.Int    // Signature component
   // R          *big.Int    // Signature component
   // S          *big.Int    // Signature component
}
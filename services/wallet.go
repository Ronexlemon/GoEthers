package services

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func GeneratePrivateKey() ([]byte, []byte,string, error) {
	// Generate a new ECDSA private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil,"", err
	}

	// Extract the public key from the private key
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	//address
	address := crypto.PubkeyToAddress(*publicKey).String()

	// Convert private and public keys to byte slices
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	return privateKeyBytes, publicKeyBytes, address, nil


	
}

package base

import (
	"fmt"
	"reflect"
	"crypto/sha256"
	"workspace/etm-go-lib/src/utils"
)

type Account struct {
	PublicKey string
}

func (a Account) IsEmpty() bool {
	return reflect.DeepEqual(a, Account{})
}

func (a Account) GetKeypairBySecret(secret string) (string, string) {
	hash := sha256.Sum256([]byte(secret))
	ed := utils.Ed{}
	keypair := ed.MakeKeypair(hash[:])
	
	publicKey := fmt.Sprintf("%x", keypair.PublicKey)
	privateKey := fmt.Sprintf("%x", keypair.PrivateKey)
	
	return publicKey, privateKey
}

func (a Account) GetAddressByPublicKey(publicKey []byte) string {
	address := utils.Address{}
	return address.GenerateAddresss(publicKey)
}

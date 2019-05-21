package base

import (
	"fmt"
	"bytes"
	"crypto/sha256"
	"workspace/etm-go-lib/src/utils"
)

type Block struct {
	Id                   string        `json:"id"`
	Version              string        `json:"version"`
	TotalAmount          int64         `json:"totalAmount"`
	TotalFee             int64         `json:"totalFee"`
	Reward               int64         `json:"reward"`
	PayloadHash          string        `json:"payloadHash"`
	Timestamp            int64         `json:"timestamp"`
	NumberOfTransactions int           `json:"NnumberOfTransactions"`
	PayloadLength        int64         `json:"payloadLength"`
	PreviousBlock        string        `json:"previousBlock"`
	GeneratorPublicKey   string        `json:"generatorPublicKey"`
	Transactions         []Transaction `json:"transactions"`
}

func (block *Block) Create(data UserData) {

}

func (block *Block) GetBytes() []byte {
	bb := bytes.NewBuffer([]byte{})
	
	return bb.Bytes()
}

func (block *Block) GetHash() [32]byte {
	_bytes := block.GetBytes()
	_hash := sha256.Sum256(_bytes)
	return _hash
}

func (block *Block) GetId() string {
	return fmt.Sprintf("%x", block.GetHash())
}

func (block *Block) GetSignature(keypair utils.Keypair) string {
	_hash := block.GetHash()
	ed := utils.Ed{}
	_sign := ed.Sign(_hash[:], keypair)
	return fmt.Sprintf("%x", _sign)
}

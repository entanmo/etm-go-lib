package base

import (
	"fmt"
	"bytes"
	"crypto/sha256"
	"workspace/etm-go-lib/src/utils"
	"reflect"
)

const maxPayloadLength  = 8 * 1024 * 1024

type Block struct {
	Id                   string        `json:"id"`
	Height               int64        `json:"height"`
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

type BlockData struct {
	Transactions  []Transaction
	PreviousBlock Block
}

func sortTransactions(trs []Transaction) []Transaction {
	return trs
}

func (block *Block) IsEmpty() bool {
	return reflect.DeepEqual(block, Block{})
}

func (block *Block) Create(data BlockData) {
	//trs := sortTransactions(data.Transactions)
	//var nextHeight int64 = 1
	//if !data.PreviousBlock.IsEmpty() {
	//	nextHeight = data.PreviousBlock.Height + 1
	//}
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

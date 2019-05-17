package base

import (
	"fmt"
	"workspace/etm-go-lib/src/utils"
	"crypto/sha256"
	"bytes"
	"encoding/binary"
)

type TrType uint8

const (
	TRANSFER     TrType = 0
	DELEGATE     TrType = 2
	UNDELEGATE   TrType = 120
	LOCK         TrType = 101
	UNLOCK       TrType = 102
	VOTE         TrType = 3
	DELAY        TrType = 110
	SECOND       TrType = 1
	MULTI        TrType = 4
	UIA_ISSYER   TrType = 9
	UIA_ASSET    TrType = 10
	UIA_FALG     TrType = 11
	UIA_ACL      TrType = 12
	UIA_ISSUE    TrType = 13
	UIA_TRANSFER TrType = 14
)

type Transaction struct {
	Type               TrType      `json:"type"`
	Id                 string      `json:"id"`
	Fee                float64     `json:"fee"`
	Amount             float64     `json:"amount"`
	Timestamp          int64       `json:"timestamp"`
	Asset              interface{} `json:"asset"`
	Args               []string    `json:"args"`
	Message            string `json:"message"`
	Signature          string      `json:"signature"`
	SignSignature      string      `json:"signSignature"`
	SenderPublicKey    string
	RequesterPublicKey string
	RecipientId        string
}

type SubTr interface {
	create(tr *Transaction, data UserData)
	getBytes(tr *Transaction) []byte
}

var trs = make(map[TrType]SubTr)

func RegisterTrs(trType TrType, tr SubTr) {
	trs[trType] = tr
}

type UserData struct {
	Type          TrType
	Amount        float64
	Fee           float64
	Timestamp     int64
	Asset         interface{}
	Args          []string
	Message       string
	Sender        Account
	Keypair       utils.Keypair
	SecondKeypair utils.Keypair
}

func (tr *Transaction) Create(data UserData) {
	if data.Sender.IsEmpty() {
		return
	}
	if data.Keypair.IsEmpty() {
		return
	}
	
	tr.Type = data.Type
	tr.Amount = 0
	tr.Fee = data.Fee
	tr.Timestamp = data.Timestamp
	tr.Asset = data.Asset
	tr.Args = data.Args
	tr.Message = data.Message
	
	trs[data.Type].create(tr, data) //构建对应子交易数据
	
	tr.Signature = tr.GetSignature(data.Keypair)
	if (data.Type != 1 && data.SecondKeypair.IsEmpty()) {
		tr.SignSignature = tr.GetSignature(data.SecondKeypair)
	}
	
	tr.Id = tr.GetId();
}

func (tr *Transaction) GetBytes(skipSignature bool, skipSecondSignature bool) []byte {
	size := 1 + 4 + 32 + 32 + 8 + 8 + 64 + 64;
	assetBytes := trs[tr.Type].getBytes(tr)
	assetSize := len(assetBytes)
	
	bb := bytes.NewBuffer(make([]byte, size+assetSize))
	
	binary.Write(bb, binary.LittleEndian, uint8(tr.Type))
	binary.Write(bb, binary.LittleEndian, uint32(tr.Timestamp))
	
	bb.WriteString(tr.SenderPublicKey)
	bb.WriteString(tr.RequesterPublicKey)
	bb.WriteString(tr.RecipientId)
	
	binary.Write(bb, binary.LittleEndian, uint64(tr.Amount))
	
	if tr.Message != "" {
		bb.WriteString(string(tr.Message))
	}
	
	if assetSize > 0 {
		bb.Write(assetBytes)
	}
	
	if !skipSignature && tr.Signature != "" {
		bb.WriteString(tr.Signature)
	}
	
	if !skipSecondSignature && tr.SignSignature != "" {
		bb.WriteString(tr.SignSignature)
	}
	
	return bb.Bytes()
}

func (tr *Transaction) GetHash() [32]byte {
	_bytes := tr.GetBytes(false, false)
	_hash := sha256.Sum256(_bytes)
	return _hash
}

func (tr *Transaction) GetId() string {
	return fmt.Sprintf("%x", tr.GetHash())
}

func (tr *Transaction) GetSignature(keypair utils.Keypair) string {
	_hash := tr.GetHash()
	_sign := utils.Ed{}.Sign(_hash[:], keypair)
	return fmt.Sprintf("%x", _sign)
}

func (tr *Transaction) GetMultiSignature(keypair utils.Keypair) string {
	_bytes := tr.GetBytes(true, true)
	_hash := sha256.Sum256(_bytes)
	_sign := utils.Ed{}.Sign(_hash[:], keypair)
	return fmt.Sprintf("%x", _sign)
}

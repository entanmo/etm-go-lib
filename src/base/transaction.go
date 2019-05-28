package base

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"workspace/etm-go-lib/src/utils"
	"encoding/binary"
	"encoding/hex"
)

type TrType uint8

const (
	TRANSFER   TrType = 0
	DELEGATE   TrType = 2
	UNDELEGATE TrType = 120
	LOCK       TrType = 101
	UNLOCK     TrType = 102
	VOTE       TrType = 3
	DELAY      TrType = 110
	SECOND     TrType = 1
	MULTI      TrType = 4
	
	UIA_ISSUER   TrType = 9
	UIA_ASSET    TrType = 10
	UIA_FALG     TrType = 11
	UIA_ACL      TrType = 12
	UIA_ISSUE    TrType = 13
	UIA_TRANSFER TrType = 14
)

type Asset struct {
	Signature   Second
	Vote        Vote
	Delegate    Delegate
	UiaIssuer   UiaIssuer
	UiaAsset    UiaAsset
	UiaFlags    UiaFlags
	UiaAcl      UiaAcl
	UiaIssue    UiaIssue
	UiaTransfer UiaTransfer
}

type Transaction struct {
	Type               TrType   `json:"type"`
	Id                 string   `json:"id"`
	Fee                int64    `json:"fee"`
	Amount             int64    `json:"amount"`
	Timestamp          int64    `json:"timestamp"`
	RecipientId        string   `json:"recipientId"`
	Asset              Asset    `json:"asset"`
	Args               []string `json:"args"`
	Message            string   `json:"message"`
	Signature          string   `json:"signature"`
	SignSignature      string   `json:"signSignature"`
	SenderPublicKey    string   `json:"senderPublicKey"`
	RequesterPublicKey string   `json:"requesterPublicKey"`
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
	Type           TrType
	Amount         int64
	Fee            int64
	Timestamp      int64
	RecipientId    string
	Asset          Asset
	Args           []string
	Message        string
	Sender         Account
	Keypair        utils.Keypair
	SecondKeypair  utils.Keypair
	Votes          []string
	Username       string
	Name           string
	Desc           string
	Maximun        string
	Precision      byte
	Strategy       string
	AllawWriteOff  byte
	AllowWhiteList byte
	AllowBlackList byte
	Currency       string
	UiaAmount      string
	FlagType       byte
	Flag           byte
	Operator       string
	List           []string
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
	tr.SenderPublicKey = data.Sender.PublicKey
	tr.Asset = data.Asset
	tr.Args = data.Args
	tr.Message = data.Message
	
	trs[data.Type].create(tr, data) //构建对应子交易数据
	
	tr.Signature = tr.GetSignature(data.Keypair)
	if data.Type != 1 && !data.SecondKeypair.IsEmpty() {
		tr.SignSignature = tr.GetSignature(data.SecondKeypair)
	}
	
	tr.Id = tr.GetId();
}

func (tr *Transaction) GetBytes(skipSignature bool, skipSecondSignature bool) []byte {
	//size := 1 + 4 + 32 + 32 + 8 + 8 + 64 + 64
	
	assetBytes := trs[tr.Type].getBytes(tr)
	assetSize := len(assetBytes)
	
	//bb := bytes.NewBuffer(make([]byte, size+assetSize))
	bb := bytes.NewBuffer([]byte{})
	
	binary.Write(bb, binary.LittleEndian, uint8(tr.Type))
	binary.Write(bb, binary.LittleEndian, uint32(tr.Timestamp))
	
	if tr.SenderPublicKey != "" {
		senderPublicKeyBytes, _ := hex.DecodeString(tr.SenderPublicKey)
		bb.Write(senderPublicKeyBytes)
	}
	
	if tr.RequesterPublicKey != "" {
		requesterPublicKeyBytes, _ := hex.DecodeString(tr.RequesterPublicKey)
		bb.Write(requesterPublicKeyBytes)
	}
	
	if tr.RecipientId != "" {
		bb.WriteString(tr.RecipientId)
	} else {
		for i := 0; i < 8; i++ {
			bb.WriteByte(0);
		}
	}
	
	binary.Write(bb, binary.LittleEndian, uint64(tr.Amount))
	
	if tr.Message != "" {
		bb.WriteString(string(tr.Message))
	}
	
	if tr.Args != nil && len(tr.Args) > 0 {
		for i := 0; i < len(tr.Args); i++ {
			bb.WriteString(tr.Args[i])
		}
	}
	
	if assetSize > 0 {
		bb.Write(assetBytes)
	}
	
	if !skipSignature && tr.Signature != "" {
		signatureBytes, _ := hex.DecodeString(tr.Signature)
		bb.Write(signatureBytes)
	}
	
	if !skipSecondSignature && tr.SignSignature != "" {
		signSignatureBytes, _ := hex.DecodeString(tr.SignSignature)
		bb.Write(signSignatureBytes)
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
	_sign := ed.Sign(_hash[:], keypair)
	return fmt.Sprintf("%x", _sign)
}

func (tr *Transaction) GetMultiSignature(keypair utils.Keypair) string {
	_bytes := tr.GetBytes(true, true)
	_hash := sha256.Sum256(_bytes)
	_sign := ed.Sign(_hash[:], keypair)
	return fmt.Sprintf("%x", _sign)
}

package base

import (
	"fmt"
	"bytes"
	"crypto/sha256"
	"workspace/etm-go-lib/src/utils"
	"reflect"
	"encoding/binary"
	"encoding/hex"
	"sort"
)

const maxPayloadLength = 8 * 1024 * 1024

var ed = utils.Ed{}
var blockStatus = utils.BlockStatus{}

type Block struct {
	Id                   string        `json:"id"`
	Height               int64         `json:"height"`
	Version              int           `json:"version"`
	TotalAmount          int64         `json:"totalAmount"`
	TotalFee             int64         `json:"totalFee"`
	Reward               int64         `json:"reward"`
	PayloadHash          string        `json:"payloadHash"`
	Timestamp            int64         `json:"timestamp"`
	NumberOfTransactions int           `json:"NnumberOfTransactions"`
	PayloadLength        int           `json:"payloadLength"`
	PreviousBlock        string        `json:"previousBlock"`
	GeneratorPublicKey   string        `json:"generatorPublicKey"`
	Transactions         []Transaction `json:"transactions"`
	BlockSignature       string        `json:"blockSignature"`
}

type BlockData struct {
	Transactions  []Transaction
	PreviousBlock Block
	Timestamp     int64
	Keypair       utils.Keypair
}

type SortTrs []Transaction

func (s SortTrs) Len() int {
	return len(s)
}
func (s SortTrs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortTrs) Less(i, j int) bool {
	if (s[i].Type != s[j].Type) {
		if (s[i].Type == 1) {
			return true
		}
		if (s[j].Type == 1) {
			return false
		}
		return s[i].Type > s[j].Type
	}
	if (s[i].Amount != s[j].Amount) {
		return s[i].Amount > s[j].Amount
	}
	return s[i].Id > s[j].Id
}

func sortTransactions(trs []Transaction) []Transaction {
	sort.Sort(SortTrs(trs))
	return trs
}

func (block *Block) IsEmpty() bool {
	return reflect.DeepEqual(block, Block{})
}

func (block *Block) Create(data BlockData) {
	trs := sortTransactions(data.Transactions)
	var nextHeight int64 = 1
	if !data.PreviousBlock.IsEmpty() {
		nextHeight = data.PreviousBlock.Height + 1
	}
	reward := blockStatus.CalcReward(nextHeight)
	var totalFee int64
	var totalAmount int64
	var size int
	
	blockTrs := trs[:]
	payloadHash := sha256.New()
	
	for i := 0; i < len(trs); i++ {
		bs := trs[i].GetBytes(false, false)
		
		if size+len(bs) > maxPayloadLength {
			blockTrs = trs[:i]
			break
		}
		
		size += len(bs)
		totalFee += trs[i].Fee
		totalAmount += trs[i].Amount
		
		payloadHash.Write(bs)
	}
	
	block.Version = 0
	block.TotalAmount = totalAmount
	block.TotalFee = totalFee
	block.Reward = reward
	block.PayloadHash = fmt.Sprintf("%x", payloadHash.Sum([]byte{}))
	block.Timestamp = data.Timestamp
	block.NumberOfTransactions = len(blockTrs)
	block.PayloadLength = size
	block.PreviousBlock = data.PreviousBlock.Id
	block.GeneratorPublicKey = fmt.Sprintf("%x", data.Keypair.PublicKey)
	block.Transactions = blockTrs
	
	block.BlockSignature = block.GetSignature(data.Keypair)
}

func (block *Block) GetBytes() []byte {
	bb := bytes.NewBuffer([]byte{})
	
	binary.Write(bb, binary.LittleEndian, uint32(block.Version))
	binary.Write(bb, binary.LittleEndian, uint32(block.Timestamp))
	
	if block.PreviousBlock != "" {
		bb.WriteString(block.PreviousBlock)
	} else {
		bb.WriteString("0")
	}
	
	binary.Write(bb, binary.LittleEndian, uint32(block.NumberOfTransactions))
	binary.Write(bb, binary.LittleEndian, uint64(block.TotalAmount))
	binary.Write(bb, binary.LittleEndian, uint64(block.TotalFee))
	binary.Write(bb, binary.LittleEndian, uint64(block.Reward))
	
	binary.Write(bb, binary.LittleEndian, uint32(block.PayloadLength))
	
	payloadHashBytes, _ := hex.DecodeString(block.PayloadHash)
	bb.Write(payloadHashBytes)
	
	generatorPublicKeyBytes, _ := hex.DecodeString(block.GeneratorPublicKey)
	bb.Write(generatorPublicKeyBytes)
	
	if block.BlockSignature != "" {
		blockSignatureBytes, _ := hex.DecodeString(block.BlockSignature)
		bb.Write(blockSignatureBytes)
	}
	
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
	_sign := ed.Sign(_hash[:], keypair)
	return fmt.Sprintf("%x", _sign)
}

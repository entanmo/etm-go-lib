package main

import (
	"fmt"
	
	"workspace/etm-go-lib/src/base"
	"workspace/etm-go-lib/src/utils"
	"crypto/sha256"
)

func generateData(secret string, data base.UserData) base.UserData {
	hash := sha256.Sum256([]byte(secret))
	sender := base.Account{}
	ed := utils.Ed{}
	keypair := ed.MakeKeypair(hash[:])
	sender.PublicKey = fmt.Sprintf("%x", keypair.PublicKey)
	//sender.PublicKey = string(keypair.PublicKey[:])
	//sender.PublicKey = hex.Dump(keypair.PublicKey)
	
	data.Keypair = keypair
	data.Sender = sender
	
	return data
}

func trTransfer() base.Transaction {
	data := base.UserData{
		Type:        0,
		Amount:      123456789,
		Fee:         10000000,
		Timestamp:   18972768,
		RecipientId: "A79wqbYgZC5Bb923wWDXKjD7KBDn5BD6gg",
	}
	data1 := generateData("race forget pause shoe trick first abuse insane hope budget river enough", data)
	_tr := base.Transaction{}
	_tr.Create(data1)
	return _tr
}

func trSecond() base.Transaction {
	hash2 := sha256.Sum256([]byte("asd"))
	ed := utils.Ed{}
	secondKeypair := ed.MakeKeypair(hash2[:])
	data := base.UserData{
		Type:          1,
		Fee:           500000000,
		Timestamp:     18972768,
		SecondKeypair: secondKeypair,
	}
	data1 := generateData("immense buffalo organ pond illegal erupt prepare arrow cliff fit abstract task", data)
	_tr := base.Transaction{}
	_tr.Create(data1)
	return _tr
}

func main() {
	tr := trTransfer()
	//tr := trSecond()
	
	fmt.Println("tr=", tr)
}

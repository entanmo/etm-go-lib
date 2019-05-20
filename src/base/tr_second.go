package base

import (
	"fmt"
	"bytes"
	"encoding/hex"
)

type Second struct {
	PublicKey string
}

func init() {
	tr := Second{}
	RegisterTrs(SECOND, &tr)
}

func (second *Second) create(tr *Transaction, data UserData) {
	tr.RecipientId = ""
	tr.Amount = 0
	tr.Asset.Signature.PublicKey = fmt.Sprintf("%x", data.SecondKeypair.PublicKey)
}

func (second *Second) getBytes(tr *Transaction) []byte {
	bb := bytes.NewBuffer([]byte{})
	if tr.Asset.Signature.PublicKey != "" {
		signaturePublicKeyBytes, _ := hex.DecodeString(tr.Asset.Signature.PublicKey)
		bb.Write(signaturePublicKeyBytes)
	}
	
	return bb.Bytes()
}

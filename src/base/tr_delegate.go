package base

import (
	"bytes"
	"strings"
)

type Delegate struct {
	Username  string
	PublicKey string
}

func init() {
	tr := Delegate{}
	RegisterTrs(DELEGATE, &tr)
}

func (delegate *Delegate) create(tr *Transaction, data UserData) {
	tr.RecipientId = "A4MFB3MaPd355ug19GYPMSakCAWKbLjDTb"
	tr.Amount = 1000 * 100000000
	tr.Asset.Delegate = Delegate{
		Username:  strings.ToLower(data.Username),
		PublicKey: data.Sender.PublicKey,
	}
}

func (delegate *Delegate) getBytes(tr *Transaction) []byte {
	bb := bytes.NewBuffer([]byte{})
	bb.WriteString(tr.Asset.Delegate.Username)
	
	return bb.Bytes()
}

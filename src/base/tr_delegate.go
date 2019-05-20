package base

import (
	"bytes"
)

type Delegate struct {
	Username string
	PublicKey string
}

func init() {
	tr := Delegate{}
	RegisterTrs(DELEGATE, &tr)
}

func (delegate *Delegate) create(tr *Transaction, data UserData) {

}

func (delegate *Delegate) getBytes(tr *Transaction) []byte {
	buf := bytes.NewBuffer([]byte{})
	return buf.Bytes()
}

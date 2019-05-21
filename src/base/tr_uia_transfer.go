package base

import (
	"bytes"
)

type UiaTransfer struct {
	Currency  string
	UiaAmount string
}

func init() {
	tr := UiaTransfer{}
	RegisterTrs(UIA_TRANSFER, &tr)
}

func (transfer *UiaTransfer) create(tr *Transaction, data UserData) {
	tr.RecipientId = data.RecipientId
	tr.Amount = 0
	tr.Asset.UiaTransfer = UiaTransfer{
		Currency:  data.Currency,
		UiaAmount: data.UiaAmount,
	}
}

func (transfer *UiaTransfer) getBytes(tr *Transaction) []byte {
	bb := bytes.NewBuffer([]byte{})
	bb.WriteString(tr.Asset.UiaTransfer.Currency)
	bb.WriteString(tr.Asset.UiaTransfer.UiaAmount)
	
	return bb.Bytes()
}

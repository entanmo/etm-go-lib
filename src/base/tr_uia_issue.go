package base

import (
	"bytes"
)

type UiaIssue struct {
	Currency  string
	UiaAmount string
}

func init() {
	tr := UiaIssue{}
	RegisterTrs(UIA_ISSUE, &tr)
}

func (issue *UiaIssue) create(tr *Transaction, data UserData) {
	tr.RecipientId = ""
	tr.Amount = 0
	tr.Asset.UiaIssue = UiaIssue{
		Currency:  data.Currency,
		UiaAmount: data.UiaAmount,
	}
}

func (issue *UiaIssue) getBytes(tr *Transaction) []byte {
	bb := bytes.NewBuffer([]byte{})
	bb.WriteString(tr.Asset.UiaIssue.Currency)
	bb.WriteString(tr.Asset.UiaIssue.UiaAmount)
	
	return bb.Bytes()
}

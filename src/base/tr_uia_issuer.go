package base

import "bytes"

type UiaIssuer struct {
	Name string
	Desc string
}

func init() {
	tr := UiaIssuer{}
	RegisterTrs(UIA_ISSUER, &tr)
}

func (issuer *UiaIssuer) create(tr *Transaction, data UserData) {
	tr.RecipientId = ""
	tr.Amount = 0
	tr.Asset.UiaIssuer = UiaIssuer{
		Name: data.Name,
		Desc: data.Desc,
	}
}

func (issuer *UiaIssuer) getBytes(tr *Transaction) []byte {
	bb := bytes.NewBuffer([]byte{})
	bb.WriteString(tr.Asset.UiaIssuer.Name)
	bb.WriteString(tr.Asset.UiaIssuer.Desc)
	
	return bb.Bytes()
}

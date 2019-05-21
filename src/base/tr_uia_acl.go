package base

import "bytes"

type UiaAcl struct {
	Currency string
	Operator string
	Flag     byte
	List     []string
}

func init() {
	tr := UiaAcl{}
	RegisterTrs(UIA_ACL, &tr)
}

func (acl *UiaAcl) create(tr *Transaction, data UserData) {
	tr.RecipientId = ""
	tr.Amount = 0
	tr.Asset.UiaAcl = UiaAcl{
		Currency: data.Currency,
		Operator: data.Operator,
		Flag:     data.Flag,
		List:     data.List,
	}
}

func (acl *UiaAcl) getBytes(tr *Transaction) []byte {
	bb := bytes.NewBuffer([]byte{})
	bb.WriteString(tr.Asset.UiaAcl.Currency)
	bb.WriteString(tr.Asset.UiaAcl.Operator)
	bb.WriteByte(tr.Asset.UiaAcl.Flag)
	for i := 0; i < len(tr.Asset.UiaAcl.List); i++ {
		bb.WriteString(tr.Asset.UiaAcl.List[i])
	}
	
	return bb.Bytes()
}

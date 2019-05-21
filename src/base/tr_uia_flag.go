package base

import "bytes"

type UiaFlags struct {
	Currency string
	FlagType byte
	Flag     byte
}

func init() {
	tr := UiaFlags{}
	RegisterTrs(UIA_FALG, &tr)
}

func (flag *UiaFlags) create(tr *Transaction, data UserData) {
	tr.RecipientId = ""
	tr.Amount = 0
	tr.Asset.UiaFlags = UiaFlags{
		Currency: data.Currency,
		FlagType: data.FlagType,
		Flag:     data.Flag,
	}
}

func (flag *UiaFlags) getBytes(tr *Transaction) []byte {
	bb := bytes.NewBuffer([]byte{})
	bb.WriteString(tr.Asset.UiaFlags.Currency)
	bb.WriteByte(tr.Asset.UiaFlags.FlagType)
	bb.WriteByte(tr.Asset.UiaFlags.Flag)
	
	return bb.Bytes()
}

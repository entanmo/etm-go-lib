package base

import (
	"bytes"
)

type UiaAsset struct {
	Name           string
	Desc           string
	Maximun        string
	Precision      byte
	Strategy       string
	AllawWriteOff  byte
	AllowWhiteList byte
	AllowBlackList byte
}

func init() {
	tr := UiaAsset{}
	RegisterTrs(UIA_ASSET, &tr)
}

func (asset *UiaAsset) create(tr *Transaction, data UserData) {
	tr.RecipientId = ""
	tr.Amount = 0
	tr.Asset.UiaAsset = UiaAsset{
		Name:           data.Name,
		Desc:           data.Desc,
		Maximun:        data.Maximun,
		Precision:      data.Precision,
		Strategy:       data.Strategy,
		AllawWriteOff:  data.AllawWriteOff,
		AllowWhiteList: data.AllowWhiteList,
		AllowBlackList: data.AllowBlackList,
	}
}

func (asset *UiaAsset) getBytes(tr *Transaction) []byte {
	bb := bytes.NewBuffer([]byte{})
	bb.WriteString(tr.Asset.UiaAsset.Name)
	bb.WriteString(tr.Asset.UiaAsset.Desc)
	bb.WriteString(tr.Asset.UiaAsset.Maximun)
	bb.WriteByte(tr.Asset.UiaAsset.Precision)
	bb.WriteString(tr.Asset.UiaAsset.Strategy)
	bb.WriteByte(tr.Asset.UiaAsset.AllawWriteOff)
	bb.WriteByte(tr.Asset.UiaAsset.AllowWhiteList)
	bb.WriteByte(tr.Asset.UiaAsset.AllowBlackList)
	
	return bb.Bytes()
}

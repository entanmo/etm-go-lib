package base

type UiaAsset struct {
	Name           string
	Desc           string
	Maximun        int64
	Precision      string
	strategy       string
	AllawWriteOff  bool
	AllowWhiteList bool
	AllowBlackList bool
}

func init() {
	tr := UiaAsset{}
	RegisterTrs(UIA_ASSET, &tr)
}

func (asset *UiaAsset) create(tr *Transaction, data UserData) {

}

func (asset *UiaAsset) getBytes(tr *Transaction) []byte {
	return nil
}

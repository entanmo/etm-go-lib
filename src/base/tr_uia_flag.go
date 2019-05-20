package base

type UiaFlags struct {
	Currency string
	FlagType string
	Flag     string
}

func init() {
	tr := UiaFlags{}
	RegisterTrs(UIA_FALG, &tr)
}

func (flag *UiaFlags) create(tr *Transaction, data UserData) {

}

func (flag *UiaFlags) getBytes(tr *Transaction) []byte {
	return nil
}

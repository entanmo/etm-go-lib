package base

type UiaTransfer struct {
	Currency string
	Amount   int64
}

func init() {
	tr := UiaTransfer{}
	RegisterTrs(UIA_TRANSFER, &tr)
}

func (transfer *UiaTransfer) create(tr *Transaction, data UserData) {

}

func (transfer *UiaTransfer) getBytes(tr *Transaction) []byte {
	return nil
}

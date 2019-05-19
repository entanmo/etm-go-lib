package base

type Transfer struct {
}

func init() {
	tr := Transfer{}
	RegisterTrs(TRANSFER, &tr)
}

func (transfer *Transfer) create(tr *Transaction, data UserData) {

}

func (transfer *Transfer) getBytes(tr *Transaction) []byte {
	return nil
}

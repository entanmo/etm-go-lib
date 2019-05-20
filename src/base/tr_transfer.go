package base

type Transfer struct {
}

func init() {
	tr := Transfer{}
	RegisterTrs(TRANSFER, &tr)
}

func (transfer *Transfer) create(tr *Transaction, data UserData) {
	tr.RecipientId = data.RecipientId
	tr.Amount = data.Amount
}

func (transfer *Transfer) getBytes(tr *Transaction) []byte {
	return nil
}

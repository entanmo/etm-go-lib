package base

type Delay struct {
}

func init() {
	tr := Delay{}
	RegisterTrs(DELAY, &tr)
}

func (delay *Delay) create(tr *Transaction, data UserData) {
	tr.RecipientId = data.RecipientId
	tr.Amount = data.Amount
	tr.Args = data.Args
}

func (delay *Delay) getBytes(tr *Transaction) []byte {
	return nil
}

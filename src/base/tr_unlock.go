package base

type Unlock struct {
}

func init() {
	tr := Unlock{}
	RegisterTrs(UNLOCK, &tr)
}

func (unlock *Unlock) create(tr *Transaction, data UserData) {
	tr.Args = data.Args
}

func (unlock *Unlock) getBytes(tr *Transaction) []byte {
	return nil
}

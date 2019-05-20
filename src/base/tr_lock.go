package base

type Lock struct {
}

func init() {
	tr := Lock{}
	RegisterTrs(LOCK, &tr)
}

func (lock *Lock) create(tr *Transaction, data UserData) {
	tr.Args = data.Args
}

func (lock *Lock) getBytes(tr *Transaction) []byte {
	return nil
}

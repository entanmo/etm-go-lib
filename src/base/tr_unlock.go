package base

type Unlock struct {
}

func init() {
	tr := Vote{}
	RegisterTrs(UNLOCK, &tr)
}

func (unlock *Unlock) create(tr *Transaction, data UserData) {

}

func (unlock *Unlock) getBytes(tr *Transaction) []byte {
	return nil
}

package base

type Second struct {
}

func init() {
	tr := Vote{}
	RegisterTrs(SECOND, &tr)
}

func (second *Second) create(tr *Transaction, data UserData) {

}

func (second *Second) getBytes(tr *Transaction) []byte {
	return nil
}

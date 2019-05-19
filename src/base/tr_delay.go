package base

type Delay struct {
}

func init() {
	tr := Vote{}
	RegisterTrs(DELAY, &tr)
}

func (delay *Delay) create(tr *Transaction, data UserData) {

}

func (delay *Delay) getBytes(tr *Transaction) []byte {
	return nil
}

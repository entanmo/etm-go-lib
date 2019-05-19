package base

type Vote struct {
}

func init() {
	tr := Vote{}
	RegisterTrs(VOTE, &tr)
}

func (vote *Vote) create(tr *Transaction, data UserData) {

}

func (vote *Vote) getBytes(tr *Transaction) []byte {
	return nil
}

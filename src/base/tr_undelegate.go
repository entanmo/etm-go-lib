package base

type Undelegate struct {
}

func init() {
	tr := Vote{}
	RegisterTrs(UNDELEGATE, &tr)
}

func (undelegate *Undelegate) create(tr *Transaction, data UserData) {

}

func (undelegate *Undelegate) getBytes(tr *Transaction) []byte {
	return nil
}

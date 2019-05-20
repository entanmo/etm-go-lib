package base

type Undelegate struct {
}

func init() {
	tr := Undelegate{}
	RegisterTrs(UNDELEGATE, &tr)
}

func (undelegate *Undelegate) create(tr *Transaction, data UserData) {

}

func (undelegate *Undelegate) getBytes(tr *Transaction) []byte {
	return nil
}

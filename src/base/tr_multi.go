package base

type Multi struct {
}

func init() {
	tr := Vote{}
	RegisterTrs(MULTI, &tr)
}

func (multi *Multi) create(tr *Transaction, data UserData) {

}

func (multi *Multi) getBytes(tr *Transaction) []byte {
	return nil
}

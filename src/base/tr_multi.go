package base

type Multi struct {
}

func init() {
	tr := Multi{}
	RegisterTrs(MULTI, &tr)
}

func (multi *Multi) create(tr *Transaction, data UserData) {

}

func (multi *Multi) getBytes(tr *Transaction) []byte {
	return nil
}

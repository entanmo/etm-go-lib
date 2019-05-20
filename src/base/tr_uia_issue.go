package base

type UiaIssue struct {
	Currency string
	Amount   int64
}

func init() {
	tr := UiaIssue{}
	RegisterTrs(UIA_ISSUE, &tr)
}

func (issue *UiaIssue) create(tr *Transaction, data UserData) {

}

func (issue *UiaIssue) getBytes(tr *Transaction) []byte {
	return nil
}

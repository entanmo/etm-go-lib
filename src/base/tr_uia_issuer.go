package base

type UiaIssuer struct {
	Name string
	Desc string
}

func init() {
	tr := UiaIssuer{}
	RegisterTrs(UIA_ISSUER, &tr)
}

func (issuer *UiaIssuer) create(tr *Transaction, data UserData) {

}

func (issuer *UiaIssuer) getBytes(tr *Transaction) []byte {
	return nil
}

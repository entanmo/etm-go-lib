package base

type UiaAcl struct {
}

func init() {
	tr := UiaAcl{}
	RegisterTrs(UIA_ACL, &tr)
}

func (acl *UiaAcl) create(tr *Transaction, data UserData) {

}

func (acl *UiaAcl) getBytes(tr *Transaction) []byte {
	return nil
}

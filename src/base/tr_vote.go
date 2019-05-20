package base

type Vote struct {
	Votes []string
}

func init() {
	tr := Vote{}
	RegisterTrs(VOTE, &tr)
}

func (vote *Vote) create(tr *Transaction, data UserData) {
	tr.Asset.Vote.Votes = data.Votes
}

func (vote *Vote) getBytes(tr *Transaction) []byte {
	return nil
}

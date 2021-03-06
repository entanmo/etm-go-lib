package base

import "bytes"

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
	bb := bytes.NewBuffer([]byte{})
	for i := 0; i < len(tr.Asset.Vote.Votes); i++ {
		bb.WriteString(tr.Asset.Vote.Votes[i])
	}
	
	return bb.Bytes()
}

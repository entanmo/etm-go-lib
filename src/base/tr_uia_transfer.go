package base

type UiaTransfer struct {
	Currency string
	Amount   int64
}

func init() {
	tr := UiaTransfer{}
	RegisterTrs(UIA_TRANSFER, &tr)
}

func (transfer *UiaTransfer) create(tr *Transaction, data UserData) {
	tr.RecipientId = data.RecipientId
	tr.Amount = 0
	tr.Asset.UiaTransfer = UiaTransfer{
		Currency: data.Currency,
		Amount:   data.Amount,
	}
}

func (transfer *UiaTransfer) getBytes(tr *Transaction) []byte {
	return nil
}

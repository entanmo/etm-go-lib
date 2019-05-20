package base

import (
	"reflect"
)

type Account struct {
	PublicKey string
}

func (a Account) IsEmpty() bool {
	return reflect.DeepEqual(a, Account{})
}

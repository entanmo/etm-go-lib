package base

import "reflect"

type Account struct {
}

func (a Account) IsEmpty() bool {
	return reflect.DeepEqual(a, Account{})
}

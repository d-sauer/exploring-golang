package user

import "github.com/d-sauer/exploring-golang/hello-interface-mock/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	return u.Doer.DoSomething(42, "Hello GoMock")
}

func (u *User) UseInOrder(n int, s string) error {
	return u.Doer.DoSomething(n, s)
}

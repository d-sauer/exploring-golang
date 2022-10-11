package doer

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks github.com/d-sauer/exploring-golang/hello-interface-mock/doer Doer

type Doer interface {
	DoSomething(int, string) error
}

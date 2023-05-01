package helpers

import "fmt"

type TransactionUtils interface {
	IsPast() bool
	IsRecurrent() bool
}

func PrintInfo(a TransactionUtils) {
	fmt.Println("Reccurent Transaction", a.IsRecurrent())
	fmt.Println("Past Transaction", a.IsPast())
}

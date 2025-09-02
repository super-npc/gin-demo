package one_test

import "fmt"

type Box[T any] struct {
	Value T
}

func (b Box[T]) GetValue() T {
	return b.Value
}

func fangXinDemo() {
	b := Box[int]{}
	b.Value = 1
	b.GetValue()
	fmt.Println(b)
}

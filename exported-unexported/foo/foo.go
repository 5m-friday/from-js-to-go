package foo

import (
	"fmt"
)

type Foo struct {
	Exported string
	private int
}

const (
	Exported = "exported"
	private = "private"
)

func ExportedFunc() {
	fmt.Println(private)
}

func NewFoo() Foo {
	return Foo{
		Exported: "blah",
		private: 1,
	}
}

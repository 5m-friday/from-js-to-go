package main

import (
	"fmt"

	"github.com/5m-friday/from-js-to-go/exported-unexported/foo"
)

func main() {
	// we cannot access private constant from outside of the package
	//fmt.Println(foo.private) // error
	fmt.Println(foo.Exported)
	foo.ExportedFunc()

	f1 := foo.NewFoo()
	fmt.Println("foo 1 object:", f1)

	// we cannot set private field from outside the package
	//f2 := foo.Foo{private: 1} // error
	f2 := foo.Foo{Exported: "blah blah"}
	fmt.Println("foo 2 object:", f2)
}

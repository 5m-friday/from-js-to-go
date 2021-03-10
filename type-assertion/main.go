package main

import (
	"errors"
	"fmt"
)

func main() {
	customPrint(1)
	customPrint("hello world")
	customPrint(errors.New("some error"))
	customPrint([]byte("hello world"))
}

func customPrint(param interface{}) {
	switch param.(type) {
	case int:
		fmt.Println("got some int")
	case string:
		fmt.Println("got some string")
	case error:
		fmt.Println("got some error")
	default:
		fmt.Println("got something else")
	}
}

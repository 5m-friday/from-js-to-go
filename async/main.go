package main

import (
	"fmt"
	"time"
)

func main() {
	go sum(2, 3)
	go sum(1, 2)
	go sum(-3, 2)
	time.Sleep(time.Second)
}

func sum(x, y int) {
	fmt.Println(x+y)
}

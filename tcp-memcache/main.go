package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

var mu sync.RWMutex

func main() {
	cache := cache{}
	listener, err := net.Listen(
		"tcp",
		":8080",
	)
	defer listener.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		// handle errors, you bustard
		conn, _ := listener.Accept()
		fmt.Fprintf(
			conn,
			"Welcome to CACHE 1.0\n-> ",
		)
		go listen(conn, cache)
	}
}

func listen(c net.Conn, mem cache) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		l := strings.ToLower(strings.TrimSpace(scanner.Text()))
		values := strings.Split(l, " ")
		switch {
		case len(values) == 3 && values[0] == "set":
			mem.Set(values[1], values[2])
			fmt.Fprintf(c, "OK\n-> ")
		case len(values) == 2 && values[0] == "get":
			fmt.Fprintf(c, "%s\n-> ",
				mem.Get(values[1]))
		case len(values) == 1 && values[0] == "exit":
			c.Close()
		default:
			fmt.Fprintf(c, "UNKNOWN: %s\n-> ", l)
		}
	}
}

type cache map[string]string

func (c cache) Set(key, value string) {
	mu.Lock()
	defer mu.Unlock()
	c[key] = value
}
func (c cache) Get(key string) string {
	mu.RLock()
	defer mu.RUnlock()
	return c[key]
}

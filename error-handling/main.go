package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	strNum := "1"
	http.HandleFunc("/cats", catsController(strNum))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("could not run http server: %v", err)
	}
}

func catsController(num string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// the error handling in the controller sucks
		intNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("could not parse number: %v", err)
			return
		}
		// anonymous structs can get very ugly
		// especially when used in conjunction with unnamed fields
		res := struct {
			Number int `json:"number"`
		}{intNum}
		bs, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("could not marshal response: %v", err)
			return
		}
		_, err = w.Write(bs)
		if err != nil {
			fmt.Printf("could not write response: %v", err)
			return
		}
	}
}

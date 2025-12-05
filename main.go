package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("=============================")
	fmt.Println("hello this is from docker PC and whats up a to the b")
	fmt.Println("=============================")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world hello world heyyy whats up"))
	})
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		fmt.Println(er)
	}

}

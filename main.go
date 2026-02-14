package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("=============================")
	fmt.Println("hello this is from docker PC and whats up a to the b")
	fmt.Println("=============================")
	fmt.Print(uuid.New().String()[:5])
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world hello world heyyy whats up"))
	// })
	// er := http.ListenAndServe(":8080", nil)
	// if er != nil {
	// 	fmt.Println(er)
	// }

}

package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "hello,http")
}
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("http serve error ")
		return
	}

}

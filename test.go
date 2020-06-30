package main

import (
	"fmt"
	"net/http"
)

func blogServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintf(w, "abcd")
}

func main() {
	fmt.Println("-----")
	http.HandleFunc("/token", blogServer)
	err := http.ListenAndServe(":8000", nil)
	fmt.Println("error", err)
	if err == nil {
		fmt.Println("OK")
	}
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
	time.Sleep(5 * time.Second)
}

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":1235", nil)
}

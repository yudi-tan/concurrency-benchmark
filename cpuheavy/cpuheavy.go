package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func serve(w http.ResponseWriter, r *http.Request) {
	tmp := fibonacci(45)
	fmt.Fprintln(w, strconv.Itoa(tmp))
}


func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":1234", nil)
}

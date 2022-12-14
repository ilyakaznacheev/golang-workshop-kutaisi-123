package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from %s\n", time.Now())
}

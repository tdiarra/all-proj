package main

import (
	"fmt"
	"net/http"
	t "time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Println(t.Now())
	})

	http.ListenAndServe(":80", nil)
}

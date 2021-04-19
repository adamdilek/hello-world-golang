package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello World From Sunny California - master!\n")

	})
	http.ListenAndServe(":80", nil)
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

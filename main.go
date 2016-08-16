package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path[1:]) > 0 {
		if strings.EqualFold(r.URL.Path[1:], "safe") {
			fmt.Fprintf(w, "DROP TABLE ")
		} else {
			fmt.Fprintf(w, "Hi here %s!", r.URL.Path[1:])
		}
	} else {
		fmt.Fprintf(w, "")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}

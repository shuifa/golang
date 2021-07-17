package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/double", doubleFunc)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func doubleFunc(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("v")
	if text == "" {
		http.Error(w, "missing error", http.StatusBadRequest)
		return
	}

	v, ok := strconv.Atoi(text)
	if ok != nil {
		http.Error(w, "not a number" + text , http.StatusBadRequest)
		return
	}

	if _, err := fmt.Fprintln(w, v * 2); err != nil {
		http.Error(w, "cannot wright response", http.StatusBadRequest)
		return
	}

	return
}
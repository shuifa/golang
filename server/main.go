package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("http req start")
	ctx := r.Context()
	select {
	case <-time.After(time.Second * 3):
		_, err := fmt.Fprintln(w, "hello golang")
		if err != nil {
			log.Fatalln(err)
		}

	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
	fmt.Println("http req end")
}

func main() {
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

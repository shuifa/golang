package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	rBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rBytes))
}

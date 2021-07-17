package main

import (
	"context"
	"fmt"
	"time"
)

func do(ctx context.Context) {
	select {
	case <- time.After(time.Second * 3):
		fmt.Println("finished doing something")
	case <- ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(time.Second * 4)
		cancel()
	}()
	do(ctx)
}

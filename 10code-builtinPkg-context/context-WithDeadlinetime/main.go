package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	cancelCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second * 5))
	
	defer cancel()
	go task(cancelCtx)

	time.Sleep(time.Second * 7)
}

func task(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("优雅退出")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}

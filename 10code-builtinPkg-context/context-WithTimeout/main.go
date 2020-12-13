package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	cancelCtx, cancel := context.WithTimeout(ctx, time.Second*3)

	defer cancel()

	go task(cancelCtx)

	time.Sleep(time.Second * 5)
	// cancelFunc()

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

package main

import (
	"context"
	"fmt"
	"time"
)


func main() {
	ctx := context.Background()

	cancelCtx, cancelFunc := context.WithCancel(ctx)

	go task(cancelCtx)
	time.Sleep(time.Second * 3)
	cancelFunc()
	time.Sleep(time.Second * 1)


}


func task(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("优雅退出")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second*1)
			i++
		}
	}
}
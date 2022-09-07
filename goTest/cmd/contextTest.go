package main

import (
	"context"
	"fmt"
	"time"
)

func contextTest(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine finished")
			return
		default:
		}
		num = num + 1
		fmt.Println(ctx.Deadline())
		fmt.Println("running", num)
	}
}
func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	i := 0
	go contextTest(ctx, i)
	time.Sleep(3 * time.Second)

}

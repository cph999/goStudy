package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	_ "github.com/panjf2000/ants"
	"time"
)

func Task() {
	fmt.Println("Hello World", "123")
	time.Sleep(5 * time.Second)
}

func main() {
	wait := make(chan struct{}, 0)
	pool, _ := ants.NewPool(10, ants.WithNonblocking(true))
	defer pool.Release()

	task := func() {
		Task()
	}

	for i := 0; i <= 10; i++ {
		pool.Submit(task)
		time.Sleep(1 * time.Second)
	}
	<-wait
}

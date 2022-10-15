package main

import "fmt"

// 当前执行的goroutine中有一个defer链表的头指针，同时有一个panic链表的头指针，每次采用头插法
// https://www.bilibili.com/video/BV155411Y7XT/?spm_id_from=333.337.search-card.all.click&vd_source=7d0bdde140a086a22ee11cfbf7607fb7
func main() {
	f()
	fmt.Println("return normally from f")
}
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from f")
		}
	}()
	g(0)
	fmt.Println("return normally from g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("panic")
		panic("i > 3")
	}
	defer fmt.Println("defer in g", i)
	fmt.Println("print in g", i)
	g(i + 1)
}

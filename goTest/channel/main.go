package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

var (
	// map[fileName]channel
	requestMap   sync.Map
	readChannel  = make(chan struct{}, 3)
	writeChannel = make(chan struct{}, 0)
	quitChannel  = make(chan struct{}, 0)
	temMap       = make(map[int]chan string, 10)
)

type Request struct {
	fileName string
	priority int
}

// filename 文件路径, priority 读取优先级，越小越大
func readFile(fileName string, priority int) {
	req := Request{fileName, priority}
	requestMap.Store(req, make(chan string, 1000))
	typeC, ok := requestMap.Load(req)
	if !ok {
		fmt.Println("It's not ok for type")
		return
	}
	channel := typeC.(chan string)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == nil {
			channel <- line
		} else {
			if err == io.EOF {
				if len(line) != 0 {
					channel <- line + "\n"
				}
				break
			}
		}
	}
	<-readChannel
	close(channel)
	if len(readChannel) == 0 {
		writeChannel <- struct{}{}
	}
}

func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	writer := bufio.NewWriter(file)
	<-writeChannel

	requestMap.Range(func(key, value interface{}) bool {
		typeC, ok := requestMap.Load(key)
		if !ok {
			fmt.Println("It's not ok for type")
		}
		channel := typeC.(chan string)
		var a = key.(Request)
		temMap[a.priority] = channel
		return true
	})

	for i := 1; i <= 3; i++ {
		ch := temMap[i]
		for {
			if data, ok := <-ch; ok {
				writer.WriteString(data)
			} else {
				break
			}
		}
	}
	//requestMap.Range(func(k, v interface{}) bool {
	//	fmt.Println("iterate:", k, v)
	//	typeC, ok := requestMap.Load(k)
	//	if !ok {
	//		fmt.Println("It's not ok for type")
	//	}
	//	channel := typeC.(chan string)
	//	for {
	//		if data, ok := <-channel; ok {
	//			writer.WriteString(data)
	//		} else {
	//			break
	//		}
	//	}
	//	return true
	//})
	if err != nil {
		return
	}
	writer.Flush()
	quitChannel <- struct{}{}
}
func main() {
	for i := 0; i < 3; i++ {
		readChannel <- struct{}{}
	}
	go readFile("data/1.txt", 1)
	go readFile("data/2.txt", 2)
	go readFile("data/3.txt", 3)

	go writeFile("data/data.txt")
	<-quitChannel
}

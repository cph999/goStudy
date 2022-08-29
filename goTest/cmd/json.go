package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIp   string
	ServerPort int
}

func main() {
	server := new(Server)
	server.ServerName = "cphzhj"
	server.ServerPort = 521
	server.ServerIp = "121.196.223.94"

	result, err := json.Marshal(server) //序列化成json字节数组
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	fmt.Println("Marshal json: ", string(result))
	resObj := deSerializable(string(result))
	fmt.Println(resObj)
}

func deSerializable(a string) interface{} {
	server := new(Server)
	err := json.Unmarshal([]byte(a), server)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return server
}

package main

import (
	"fmt"
	"net/rpc"
)

type Watcher int

func (w *Watcher) GetInfo(arg int, result *int) error {
	*result = 1
	return nil
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("链接rpc服务器失败:", err)
	}
	var reply int
	err = client.Call("Watcher.GetInfo", 1, &reply)
	if err != nil {
		fmt.Println("调用远程服务失败", err)
	}
	fmt.Println("远程服务返回结果：", reply)
}

package main

import (
	"fmt"
	"runtime"
	"gosrpc"
)

func main() {
	//查看cpu核数
	fmt.Println("cpu核数",runtime.GOMAXPROCS(runtime.NumCPU()))
	gosrpc.Rpgo()
	gosrpc.Tcprpc()
	gosrpc.Jsonrpc()
}
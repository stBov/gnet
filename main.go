package main

import (
	"fmt"
	"runtime"
	"github.com/stBov/gnet/src/gosrpc"
)

func main() {
	//查看cpu核数 runtimes
	fmt.Println("cpu核数",runtime.GOMAXPROCS(runtime.NumCPU()))
	gosrpc.Rpgo()
	gosrpc.Tcprpc()
	gosrpc.Jsonrpc()
}
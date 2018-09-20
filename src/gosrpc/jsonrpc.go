package gosrpc

import (
	"fmt"
	"net/rpc"
	"net"
	"net/rpc/jsonrpc"
	"log"
)

func services(){
	fmt.Println("JSONrpcService START")
	rect := new(Rect);
	//注册rpc服务
	rpc.Register(rect);
	//获取tcpaddr
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8082");
	chkError(err);
	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr);
	chkError(err2);
	for {
		conn, err3 := tcplisten.Accept();
		if err3 != nil {
			continue;
		}
		//使用goroutine单独处理rpc连接请求
		//这里使用jsonrpc进行处理
		go jsonrpc.ServeConn(conn);
	}
}
func clients()  {
	fmt.Println("JSONrpcClient START")
	//连接远程rpc服务
	//这里使用jsonrpc.Dial
	rpc, err := jsonrpc.Dial("tcp", "127.0.0.1:8082");
	if err != nil {
		log.Fatal(err);
	}
	ret := 0;
	//调用远程方法
	//注意第三个参数是指针类型
	err2 := rpc.Call("Rect.Area", Params{50, 100}, &ret);
	if err2 != nil {
		log.Fatal(err2);
	}
	fmt.Println(ret);
	err3 := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret);
	if err3 != nil {
		log.Fatal(err3);
	}
	fmt.Println(ret);
	fmt.Println("JSONrpcClient END")
}
func Jsonrpc()  {
	go services()
	clients()
}

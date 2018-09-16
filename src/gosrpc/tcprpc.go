package gosrpc

import (
	"net"
	"net/rpc"
	"log"
	"fmt"
)

func chkError(err error) {
	if err != nil {
		log.Fatal(err);
	}
}

func service() {
	rect := new(Rect);
	//注册rpc服务
	rpc.Register(rect);
	//获取tcpaddr
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8081");
	chkError(err);
	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr);
	chkError(err2);
	//死循环处理连接请求
	for {
		conn, err3 := tcplisten.Accept();
		if err3 != nil {
			continue;
		}
		//使用goroutine单独处理rpc连接请求
		go rpc.ServeConn(conn);
	}
}

func client() {
	fmt.Println("TCPrpcClient START")
	//连接远程rpc服务
	//这里使用Dial，http方式使用DialHTTP，其他代码都一样
	rpc, err := rpc.Dial("tcp", "127.0.0.1:8081");
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
	fmt.Println("TCPrpcClient END")
}

func Tcprpc(){
	go service()
	client()
}
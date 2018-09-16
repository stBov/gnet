package gosrpc

import (
	"fmt"
	"net/rpc"
	"net/http"
	"log"
)

func startRpcServer() {
	fmt.Println("HTTPrpcService START")
	rect := new(Rect);
	//注册一个rect服务
	rpc.Register(rect);
	//把服务处理绑定到http协议上
	rpc.HandleHTTP();
	err := http.ListenAndServe(":8080", nil);
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println("HTTPrpcService END")
}

func tryRpcClient() {
	fmt.Println("HTTPrpcClient START")
	//连接远程rpc服务
	rpc, err := rpc.DialHTTP("tcp", "127.0.0.1:8080");
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
	fmt.Println("HTTPrpcClient END")
}

func Rpgo()  {
	go startRpcServer()
	tryRpcClient()
}


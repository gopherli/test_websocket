package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"test_websocket/pb"
)

var (
	addr = "127.0.0.1:1998"
)

var Conn pb.GreetClient

type GreetClient struct {
	cc pb.GreetClient
}

func NewGreetClientManager() GreetClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("[DailServe] Dial err", err.Error())
		return GreetClient{}
	}
	// defer conn.Close()后作用在当前方法内，关闭后conn无法传递下去
	fmt.Println("客户端已连接～")
	return GreetClient{
		cc: pb.NewGreetClient(conn),
	}
}

func StartClient() {
	cc := NewGreetClientManager()
	cc.HelloWorld()
}

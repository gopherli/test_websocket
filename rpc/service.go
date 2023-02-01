package rpc

import (
	"context"
	"fmt"
	"test_websocket/pb"
)

func (g GreetClient) HelloWorld() {
	rsp, err := g.cc.HelloWorld(context.Background(), &pb.Request{})
	if err != nil {
		fmt.Println("[HelloWorld] err", err.Error())
		return
	}
	fmt.Println("[HelloWorld] ok", rsp)
}

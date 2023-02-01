package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
)

var (
	addr = "ws://127.0.0.1:1997/ws"
)

type wsClientManager struct {
	conn        *websocket.Conn
	addr        string
	sendMsgChan chan string
	recvMsgChan chan string
	isAlive     bool
	timeout     int
}

func NewWsClientManager(addr string, timeout int) *wsClientManager {
	var conn *websocket.Conn
	sendMsgChan := make(chan string, 10)
	recvMsgChan := make(chan string, 10)
	return &wsClientManager{
		conn:        conn,
		addr:        addr,
		sendMsgChan: sendMsgChan,
		recvMsgChan: recvMsgChan,
		isAlive:     false,
		timeout:     10,
	}
}

func (ws *wsClientManager) Dail() {
	var err error
	ws.conn, _, err = websocket.DefaultDialer.Dial(ws.addr, nil)
	if err != nil {
		fmt.Println("[ConnWebSocketServe] Dial err", err.Error())
		return
	}
	ws.isAlive = true
	fmt.Println("[Dail] 客户端已连接")
}

func (ws *wsClientManager) ClientSendMsg() {
	go func() {
		for {
			msg := <-ws.sendMsgChan
			fmt.Println("发送消息", msg)
			err := ws.conn.WriteMessage(1, []byte(msg))
			if err != nil {
				fmt.Println("发送消息失败")
				continue
			}
		}
	}()
}

func (ws *wsClientManager) ClientReadMsg() {
	go func() {
		for {
			if ws.conn != nil {
				_, msg, err := ws.conn.ReadMessage()
				if err != nil {
					fmt.Println("接受消息 err", err.Error())
					ws.isAlive = false
					break
				}
				ws.recvMsgChan <- string(msg)
			}
		}
	}()
}

func (ws *wsClientManager) Msg() {
	go func() {
		a := 0
		for {
			ws.sendMsgChan <- strconv.Itoa(a)
			a++
		}
	}()
}

func (ws *wsClientManager) Recv() {
	go func() {
		for {
			msg, ok := <-ws.recvMsgChan
			if ok {
				fmt.Println("收到消息:", msg)
			}
		}
	}()
}
func (ws *wsClientManager) start() {
	for {
		if ws.conn == nil {
			ws.Dail()
			ws.ClientSendMsg()
			ws.ClientReadMsg()
			ws.Msg()
			ws.Recv()
		}
		time.Sleep(time.Second * time.Duration(ws.timeout))
	}
}

func StartWebsocketClient() {
	ws := NewWsClientManager(addr, 10)
	ws.start()
}

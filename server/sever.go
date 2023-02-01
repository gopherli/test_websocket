package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	host     = "127.0.0.1"
	port     = "1997"
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func serveWebSocket(w http.ResponseWriter, r *http.Request) {
	r.Header.Del("Origin")

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("[StartWebSocketServer] Upgrade err", err.Error())
		return
	}
	defer ws.Close()
	for {
		mty, p, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("[StartWebSocketServer] ReadMessage err", err.Error())
			return
		}

		fmt.Println(string(p))

		if err = ws.WriteMessage(mty, []byte("你好客户端～")); err != nil {
			fmt.Println("[StartWebSocketServer] WriteMessage ok")
		}
		time.Sleep(time.Second)
	}
}

func StartWebsocketServe() {
	addr := host + ":" + port
	http.HandleFunc("/ws", serveWebSocket)
	fmt.Println("[StartWebsocketServe] websocket 服务器已启动～")
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println("[StartWebsocketServe] ListenAndServe err", err.Error())
	}
}

package websocketclient

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type WebSocketClient struct {
	client *websocket.Conn
}

func (socket *WebSocketClient) Connect(uri string) *WebSocketClient {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("connect:", err)
	}
	socket.client = c
	return &WebSocketClient{}
}

func (socket *WebSocketClient) recv() string {
	_, message, err := socket.client.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return ""
	}
	return string(message)
}

func (socket *WebSocketClient) push(data string) {
	err := socket.client.WriteMessage(websocket.TextMessage, []byte(data))
	if err != nil {
		log.Println("write:", err)
		return
	}
}

func (socket *WebSocketClient) close() {
	err := socket.client.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("close:", err)
		return
	}
}
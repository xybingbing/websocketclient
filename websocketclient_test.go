package websocketclient

import "testing"

func TestClient(t *testing.T) {
	client := WebSocketClient{}
	client.Connect("ws://localhost:8080/echo")
	client.push("hello world")
	var msg = client.recv()
	t.Log(msg)
}
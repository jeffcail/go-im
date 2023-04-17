package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	Websocket *websocket.Conn
)

func WsUpgrader(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	websocket, err := upgrader.Upgrade(w, r, nil)
	return websocket, err
}

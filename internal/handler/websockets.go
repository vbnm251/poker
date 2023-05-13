package handler

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	Error:           ErrorResponse,
	ReadBufferSize:  1 << 10,
	WriteBufferSize: 1 << 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) WebsocketsEndpoint(w http.ResponseWriter, r *http.Request) {

	func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/game/index.html")
	}(w, r)

	header := http.Header{}
	header.Set("Upgrade", "websocket")

	//double set header call
	//TODO : fix
	_, err := upgrader.Upgrade(w, r, header)
	if err != nil {
		return
	}

}

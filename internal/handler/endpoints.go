package handler

import (
	"net/http"
)

func (h *Handler) MainPageEndpoint(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func (h *Handler) ConnectRandomGameEndpoint(w http.ResponseWriter, r *http.Request) {

}

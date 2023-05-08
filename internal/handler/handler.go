package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"poker/internal/logic"
)

const PathToProject = "/home/vbnm251/user/coding/backendPractice/poker/"

type Handler struct {
	Games map[string]*logic.Game
}

func NewHandler() *Handler {
	return &Handler{
		Games: make(map[string]*logic.Game),
	}
}

func (h *Handler) InitEndpoints() *mux.Router {
	r := mux.NewRouter()

	staticHandler := http.StripPrefix("/", http.FileServer(http.Dir(PathToProject+"pages/home")))
	r.PathPrefix("/").Handler(staticHandler)

	r.HandleFunc("/random", h.ConnectRandomGameEndpoint)
	r.HandleFunc("/{id}", h.WebsocketsEndpoint)

	return r
}

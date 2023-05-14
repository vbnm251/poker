package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"poker/internal/logic"
)

type Handler struct {
	Addr  string
	Games map[string]*logic.Game
}

func NewHandler(addr string) *Handler {
	return &Handler{
		Games: make(map[string]*logic.Game),
		Addr:  addr,
	}
}

func (h *Handler) InitEndpoints() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/random", h.ConnectRandomGameEndpoint)
	r.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/game/index.html")
	}).Queries("id", "{id}")
	r.HandleFunc("/ws", h.WebsocketsEndpoint).Queries("id", "{id}")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/main/index.html")
	})

	//FILE HANDLERS
	pages := r.PathPrefix("/pages").Subrouter()

	pages.HandleFunc("/main/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/main/script.js")
	})

	pages.HandleFunc("/main/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/main/style.css")
	})

	pages.HandleFunc("/game/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/game/style.css")
	})

	pages.HandleFunc("/game/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/game/script.js")
	})

	pages.HandleFunc("/game/poker-chip.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/game/poker-chip.png")
	})

	return r
}

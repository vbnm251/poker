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

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/main/index.html")
	})

	r.HandleFunc("/pages/main/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/main/script.js")
	})

	r.HandleFunc("/pages/main/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/main/style.css")
	})

	r.HandleFunc("/pages/game/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/game/style.css")
	})

	r.HandleFunc("/pages/game/poker-chip.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pages/game/poker-chip.png")
	})

	r.HandleFunc("/random", h.ConnectRandomGameEndpoint)
	r.HandleFunc("/connect", h.WebsocketsEndpoint).Queries("id", "{id}")

	return r
}

package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"poker/internal/logic"
)

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

	//USER SEE IT
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, "pages/main/index.html")
	})
	r.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, "pages/game/index.html")
	}).Queries("id", "{id}")

	//API
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/random", h.ConnectRandomGameEndpoint)
	api.HandleFunc("/ws", h.WebsocketsEndpoint).Queries("id", "{id}")
	api.HandleFunc("/gameInfo", h.GetGameInfoHandler).Queries("id", "{id}")

	//FILE HANDLERS
	pages := r.PathPrefix("/pages").Subrouter()
	pages.HandleFunc("/main/script.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, "pages/main/script.js")
	})
	pages.HandleFunc("/main/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, "pages/main/style.css")
	})
	pages.HandleFunc("/game/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, "pages/game/style.css")
	})
	pages.HandleFunc("/game/script.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, "pages/game/script.js")
	})
	pages.HandleFunc("/game/websockets.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, "pages/game/websockets.js")
	})
	pages.HandleFunc("/game/game.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, "pages/game/game.js")
	})
	pages.HandleFunc("/game/poker-chip.png", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, "pages/game/poker-chip.png")
	})

	return r
}

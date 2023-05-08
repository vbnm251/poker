package poker

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) StartServer(port string, handler http.Handler) error {
	s.server = &http.Server{
		Addr:              port,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1mb
	}

	log.Println("Server started")

	return s.server.ListenAndServe()
}

package main

import (
	"log"
	"poker"
	"poker/internal/handler"
)

func main() {
	//todo : add configuration

	s := new(poker.Server)
	h := handler.NewHandler()

	if err := s.StartServer(":8080", h.InitEndpoints()); err != nil {
		log.Fatal(err.Error())
	}
}

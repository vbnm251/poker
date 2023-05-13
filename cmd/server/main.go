package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"poker"
	"poker/internal/handler"
)

type Config struct {
	Addr string `toml:"addr"`
}

func main() {
	var cfg Config
	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		log.Fatal(err)
	}

	s := new(poker.Server)
	h := handler.NewHandler(cfg.Addr)

	if err := s.StartServer(cfg.Addr, h.InitEndpoints()); err != nil {
		log.Fatal(err.Error())
	}
}

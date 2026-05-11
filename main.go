package main

import (
	"embed"
	"log"
	"net"
	"shepherd/internal/server"
)

//go:embed public/*
var embedFS embed.FS

func main() {
	cfg := new(server.Config)
	cfg.Addr = net.JoinHostPort("", "9090")
	cfg.FS = embedFS
	log.Fatal(server.Start(cfg))
}

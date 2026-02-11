package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/kakeru-lab/rpi-edge-agent/internal/agent"
	"github.com/kakeru-lab/rpi-edge-agent/internal/api"
	"github.com/kakeru-lab/rpi-edge-agent/internal/config"
	"github.com/kakeru-lab/rpi-edge-agent/internal/memory"
)

func main() {
	cfgPath := flag.String("config", "./configs/config.example.yaml", "config path")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	store, err := memory.Open(cfg.Memory.SQLitePath)
	if err != nil {
		log.Fatalf("sqlite open error: %v", err)
	}
	defer store.Close()

	ag := agent.New(store)
	handlers := api.NewHandlers(ag)
	router := api.Router(handlers)

	srv := &http.Server{
		Addr:              cfg.Server.Addr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("starting rpi-edge-agent on %s", cfg.Server.Addr)
	log.Fatal(srv.ListenAndServe())
}

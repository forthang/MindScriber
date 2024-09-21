package main

import (
	"fmt"
	"github.com/forthang/esketit/http"
	"github.com/forthang/esketit/internal/config"
	"github.com/forthang/esketit/storage"
	"log"
)

func main() {
	cfg := config.GetConfig()

	// addr := flag.String("addr", ":8080", "address for http server")
	addr := fmt.Sprintf("%v:%v", cfg.Listen.BindIP, cfg.Listen.Port)

	s := storage.NewStorage(*cfg)

	log.Printf("Starting server on %s", addr)
	if err := http.CreateAndRunServer(s, addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

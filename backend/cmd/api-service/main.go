package api_service

import (
	"flag"
	"github.com/forthang/esketit/http"
	"github.com/forthang/esketit/storage"
	"log"
)

func main() {
	addr := flag.String("addr", ":8080", "address for http server")

	s := storage.NewStorage()

	log.Printf("Starting server on %s", *addr)
	if err := http.CreateAndRunServer(s, *addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

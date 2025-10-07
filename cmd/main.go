package main

import (
	"log"
	"net/http"
	"os"
	"posts_commets_service/cmd/service_provider"
)

func main() {
	sp := service_provider.NewServiceProvider(os.Getenv("STORAGE_MODE"))

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Println("API server listening on", addr, "storage:", os.Getenv("STORAGE_MODE"))
	if err := http.ListenAndServe(addr, sp.GetHTTPServer()); err != nil {
		log.Fatal(err)
	}
}

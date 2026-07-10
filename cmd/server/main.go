package main

import (
	"log"
	"net/http"

	"agent/internal/mcp"
)

func main() {
	mcp_server := mcp.NewMCPServer()

	log.Printf("Server started at :8080/mcp\n")
	if err := http.ListenAndServe(":8080", mcp_server); err != nil {
		log.Fatalf("Failed to ListenAndServe: %s", err)
	}
	log.Printf("Quit!\n")
}

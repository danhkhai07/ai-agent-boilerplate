package api

import (
	"agent/internal/mcp"
	"net/http"
)

func addRoutes(server *Server) {
	mux := http.NewServeMux()

	mux.Handle("/mcp", mcp.NewMCPServer())
	
	server.httpServer.Handler = mux
}

package api

import (
	// "agent/internal/mcp"
	"net/http"
)



func addRoutes(server *Server) {
	mux := http.NewServeMux()

	// mux.Handle("/mcp", mcp.NewMCPServer())
	mux.HandleFunc("GET /c", 		server.GetAllSessions)
	mux.HandleFunc("GET /c/{id}",	server.GetSession)
	mux.HandleFunc("POST /c", 		server.PostNewSession)
	mux.HandleFunc("POST /c/{id}", 	server.PostMessage)
	
	server.httpServer.Handler = mux
}

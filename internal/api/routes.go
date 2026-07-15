package api

import (
	"agent/internal/mcp"
	"net/http"
)

func addRoutes(server *Server) {
	mux := http.NewServeMux()

	// Assume running from binary in build/
	fs := http.FileServer(http.Dir("../static"))
	mux.Handle("/", fs)

	mux.Handle("/mcp", mcp.NewMCPServer())

	mux.HandleFunc("GET /c", 			server.GetAllSessions)
	mux.HandleFunc("GET /c/{id}",		server.GetSession)
	mux.HandleFunc("POST /c", 			server.PostNewSession)
	mux.HandleFunc("POST /c/{id}", 		server.PostMessage)
	mux.HandleFunc("DELETE /c/{id}",	server.DeleteSession)
	
	server.httpServer.Handler = mux
}

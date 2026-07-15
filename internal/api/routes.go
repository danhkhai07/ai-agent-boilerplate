package api

import (
	"agent/internal/mcp"
	"net/http"
	"os"
	"path/filepath"
)

func resolveStaticDir() string {
	candidates := make([]string, 0, 5)
	if configuredDir := os.Getenv("AGENT_STATIC_DIR"); configuredDir != "" {
		candidates = append(candidates, configuredDir)
	}
	candidates = append(candidates, "static", filepath.Join("..", "static"))

	if executable, err := os.Executable(); err == nil {
		executableDir := filepath.Dir(executable)
		candidates = append(candidates,
			filepath.Join(executableDir, "static"),
			filepath.Join(executableDir, "..", "static"),
		)
	}

	for _, candidate := range candidates {
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate
		}
	}
	return "static"
}

func addRoutes(server *Server) {
	mux := http.NewServeMux()

<<<<<<< HEAD
	// Serve the generated web interface both at / and under the legacy /static path.
	fs := http.FileServer(http.Dir(resolveStaticDir()))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
=======
	// Assume running from binary in build/
	fs := http.FileServer(http.Dir("../static"))
	mux.Handle("/", fs)
>>>>>>> boilerplate

	mux.Handle("/mcp", mcp.NewMCPServer())
	mux.HandleFunc("GET /c", server.GetAllSessions)
	mux.HandleFunc("GET /c/{id}", server.GetSession)
	mux.HandleFunc("POST /c", server.PostNewSession)
	mux.HandleFunc("POST /c/{id}", server.PostMessage)
	mux.HandleFunc("DELETE /c/{id}", server.DeleteSession)
	mux.Handle("GET /{$}", fs)
	mux.Handle("GET /assets/", fs)

<<<<<<< HEAD
=======
	mux.HandleFunc("GET /c", 			server.GetAllSessions)
	mux.HandleFunc("GET /c/{id}",		server.GetSession)
	mux.HandleFunc("POST /c", 			server.PostNewSession)
	mux.HandleFunc("POST /c/{id}", 		server.PostMessage)
	mux.HandleFunc("DELETE /c/{id}",	server.DeleteSession)
	
>>>>>>> boilerplate
	server.httpServer.Handler = mux
}

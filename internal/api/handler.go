package api

import (
	"agent/internal/api/dto"
	"agent/internal/domain"
	"bytes"
	"encoding/json"
	"net/http"
)

func readJSON(w http.ResponseWriter, r *http.Request, v any) bool {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, v any, status int) bool {
	bytes := new(bytes.Buffer)
	err := json.NewEncoder(bytes).Encode(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes.Bytes())
	return true
}

func (svr *Server) GetAllSessions(w http.ResponseWriter, r *http.Request) {
	sessions := svr.sessionStore.GetAll()
	resp := dto.NewGetAllSessionsResponse(sessions)
	writeJSON(w, resp, http.StatusOK)
}

func (svr *Server) GetSession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("id")
	if sessionID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session, err := svr.sessionStore.GetByID(sessionID)
	if err != nil {
		errorResponse := dto.NewErrorResponse(err.Error())
		writeJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	resp := dto.NewGetSessionReponse(session)
	writeJSON(w, resp, http.StatusOK)
}

func (svr *Server) PostMessage(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("id")
	if sessionID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var message string
	if !readJSON(w, r, &message) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	session, err := svr.sessionStore.GetByID(sessionID)
	if err != nil {
		errorResponse := dto.NewErrorResponse(err.Error())
		writeJSON(w, errorResponse, http.StatusBadRequest)
		return
	}

	session.Context.Messages = append(
		session.Context.Messages, 
		domain.Message{
			Role: domain.UserRole,
			Content: message,
		},
	)
	w.WriteHeader(http.StatusOK)
}

func (svr *Server) PostNewSession(w http.ResponseWriter, r *http.Request) {
	sessionID, err := svr.sessionStore.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := dto.NewPostNewSessionResponse(sessionID)
	writeJSON(w, resp, http.StatusCreated)
}

func (svr *Server) DeleteSession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("id")
	svr.sessionStore.DeleteByID(sessionID)
	w.WriteHeader(http.StatusNoContent)
}

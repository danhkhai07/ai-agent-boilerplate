package dto

type PostNewSessionResponse struct {
	SessionID string `json:"session_id"`
}

func NewPostNewSessionResponse(sessionID string) *PostNewSessionResponse {
	return &PostNewSessionResponse{
		SessionID: sessionID,
	}
}

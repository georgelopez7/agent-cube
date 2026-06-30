package http

import "net/http"

// WebsocketHandler - creates a new websocket connection.
func (s *Server) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	s.ws.HandleConnection(w, r)
}

package http

import "net/http"

func (s *Server) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	s.ws.HandleConnection(w, r)
}

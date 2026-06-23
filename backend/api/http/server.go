package http

import (
	"fmt"
	"log/slog"
	"net/http"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humatest"
)

type Server struct {
	Port    string
	BaseURL string
	Router  *http.ServeMux
	API     huma.API
	svc     Service
}

func NewServer(name string, version string, port string, svc Service) *Server {
	router := http.NewServeMux()
	config := huma.DefaultConfig(name, version)

	baseURL := fmt.Sprintf("http://localhost:%s", port)
	config.Servers = []*huma.Server{{URL: baseURL}}

	api := humago.New(router, config)

	return &Server{
		Port:    port,
		BaseURL: baseURL,
		Router:  router,
		API:     api,
		svc:     svc,
	}
}

func (s *Server) AddRoutes(api huma.API) {
	s.addRoutes(s.API)
}

func (s *Server) Start() {
	s.addRoutes(s.API)

	port := fmt.Sprintf(":%s", s.Port)
	fmt.Printf("Server listening on port %s\n", port)

	if err := http.ListenAndServe(port, s.Router); err != nil {
		slog.Error("Server failed", "error", err)
	}
}

func (s *Server) Mock(t *testing.T) *Server {
	_, api := humatest.New(t)
	s.addRoutes(api)
	return &Server{API: api, svc: s.svc}
}

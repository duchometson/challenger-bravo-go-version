package httpserver

import (
	app "bravo/application/currency"
	"context"
	"fmt"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *nethttp.Server
	router     *gin.Engine
}

func New(port int) *Server {
	router := gin.New()

	return &Server{
		httpServer: &nethttp.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		router: router,
	}
}

func (s *Server) AddRoute(method, route string, handler gin.HandlerFunc) {
	s.router.Handle(method, route, handler)
}

func (s *Server) ListenAndServe(currency *app.Currency) error {
	s.AddRoute(nethttp.MethodGet, "/currency", currency.Get)
	s.AddRoute(nethttp.MethodPost, "/currency", currency.Add)
	s.AddRoute(nethttp.MethodDelete, "/currency", currency.Delete)
	s.AddRoute(nethttp.MethodGet, "/convert", currency.Convert)

	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

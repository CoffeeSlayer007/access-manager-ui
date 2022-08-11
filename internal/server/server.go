package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/roppenlabs/access-manager-ui/internal/config"
	"github.com/roppenlabs/access-manager-ui/pkg/logger"
)

type Server struct {
	s *http.Server
	c *config.Config
}

func New(c *config.Config) *Server {
	host := fmt.Sprintf("%s:%d", c.Listener.Host, c.Listener.Port)

	h := getHandlers()

	logger.Info(logger.Format{Message: fmt.Sprintf("server configured with host=[%s]", host)})
	srv := &http.Server{
		Addr:    host,
		Handler: h,
	}
	return &Server{
		s: srv,
		c: c,
	}
}

func (s *Server) Run(ctx context.Context) {
	s.s.ListenAndServe()
	<-ctx.Done()
	s.s.Shutdown(ctx)
}

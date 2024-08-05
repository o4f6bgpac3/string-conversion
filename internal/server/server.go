package server

import (
	"fmt"
	"github.com/o4f6bgpac3/string-conversion/app/shared"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/o4f6bgpac3/string-conversion/gen/shared/sharedconnect"
	"github.com/o4f6bgpac3/string-conversion/internal/config"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Start() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	conversionService := shared.NewService()
	path, handler := sharedconnect.NewStringConversionServiceHandler(conversionService)
	r.Handle(path+"*", handler)

	r.Handle("/*", http.FileServer(http.Dir(filepath.Join("frontend", "dist"))))

	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	fmt.Printf("Server starting on %s\n", addr)

	return http.ListenAndServeTLS(
		addr,
		s.cfg.TLSCertFile,
		s.cfg.TLSKeyFile,
		h2c.NewHandler(r, &http2.Server{}),
	)
}

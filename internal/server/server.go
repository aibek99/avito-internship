package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"avito-internship/internal/config"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	config *config.Config
	chi    *chi.Mux
}

// NewServer is
func NewServer(
	cfg *config.Config,
) *Server {

	return &Server{config: cfg, chi: chi.NewRouter()}
}

// Run is
func (s *Server) Run(ctx context.Context) error {
	httpSrv := &http.Server{
		Addr:              s.config.ServerAddress,
		Handler:           s.chi,
		ReadHeaderTimeout: 5 * time.Second,
	}

	err := s.mapHandlers()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("[server][Run] HTTP server is started on %s\n", s.config.ServerAddress)
		if err := httpSrv.ListenAndServe(); err != nil {
			log.Printf("failed to start gateway router: %v", err)
		}
	}()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Shutdown the HTTP server
	if err := httpSrv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}
	log.Println("[server][Run] Secure server shutdown gracefully")

	wg.Wait()

	return nil
}

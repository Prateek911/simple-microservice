package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string) (*Server, error) {
	mux := setupRoutes()

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return &Server{httpServer: httpServer}, nil
}

func (s *Server) Start() error {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen :%s\n", err)
		}
		log.Println("Server up on port:", s.httpServer.Addr)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	log.Println("Shutting down server....Reason :", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown :", err)
		return err
	}

	log.Println("server exiting")
	return nil
}

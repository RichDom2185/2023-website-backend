package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Adapted from
// https://www.sobyte.net/post/2021-10/go-http-server-shudown-done-right/

type Server struct {
	Server           *http.Server
	shutdownFinished chan struct{}
}

func (s *Server) ListenAndServe() (err error) {
	if s.shutdownFinished == nil {
		s.shutdownFinished = make(chan struct{})
	}

	err = s.Server.ListenAndServe()
	if err == http.ErrServerClosed {
		// Expected error after calling Server.Shutdown().
		err = nil
	} else if err != nil {
		err = fmt.Errorf("Unexpected error from ListenAndServe: %w", err)
		return
	}

	log.Println("Waiting for shutdown finishing...")
	<-s.shutdownFinished
	log.Println("Shutdown finished")

	return
}

func (s *Server) WaitForExitingSignal(parent context.Context) {
	var waiter = make(chan os.Signal, 1) // buffered channel
	signal.Notify(waiter, syscall.SIGTERM, syscall.SIGINT)

	// blocks here until there's a signal
	<-waiter

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := s.Server.Shutdown(ctx)
	if err != nil {
		log.Println("Shutting down:", err.Error())
	} else {
		log.Println("Shutdown processed successfully")
		close(s.shutdownFinished)
	}
}

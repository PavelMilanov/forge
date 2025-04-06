package main

import (
	"context"
	"net/http"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         "0.0.0.0:5050",
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	logrus.Infof("Сервер запущен: %+v | Версия go: %s", s.httpServer.Addr, runtime.Version())
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(grace time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), grace*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		logrus.Info("Сервер остановлен")
	}
	return s.httpServer.Shutdown(ctx)
}

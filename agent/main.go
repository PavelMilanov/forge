package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PavelMilanov/forge-agent/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:00",
	})
	handler := handlers.NewHandler()
	router := handler.InitRouters()
	s := http.Server{
		Addr:         ":1323",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	logrus.Info("server started")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := router.Shutdown(ctx); err != nil {
		logrus.Fatal(err)
	}
}

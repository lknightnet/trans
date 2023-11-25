package app

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"transcription/config"
	"transcription/internal/controller/http"
	"transcription/internal/service"
	"transcription/pkg/server"
)

func Run(cfg *config.Config) {
	deps := &service.Dependencies{
		AS: cfg.AudioStorage,
	}
	services := service.NewServices(deps)
	rout := mux.NewRouter()
	http.NewTranscriptionRoutes(rout, services)
	srv := server.NewServer(rout, server.Port(cfg.HTTP.Port), server.ReadTimeout(cfg.ReadTimeout),
		server.WriteTimeout(cfg.WriteTimeout), server.ShutdownTimeout(cfg.ShutdownTimeout))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("Run: " + s.String())
	case err := <-srv.Notify():
		log.Println(errors.Wrap(err, "Run: signal.Notify"))
	}

	err := srv.Shutdown()
	if err != nil {
		log.Println(errors.Wrap(err, "Run: server shutdown"))
	}
}

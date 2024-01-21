package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"

	"github.com/bondhan/hashmap/api"
)

func main() {

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/", api.Root)
	r.Get("/init", api.Init)
	r.Get("/put", api.Put)
	r.Get("/get", api.Get)
	r.Get("/remove", api.Remove)

	RunHttpServer(logger, os.Getenv("PORT"), r)
}

func RunHttpServer(logger *logrus.Logger, port string, r *chi.Mux) {
	server := &http.Server{Addr: ":" + port, Handler: r}
	go func() {
		logger.Info("application started at port:", port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Error(err)
			return
		}
	}()

	// Setting up a channel to capture system signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	//wait forever until 1 of signals above are received
	<-stop

	// send warning that we are closing
	logger.Warnf("got signal: %v, closing any connection gracefully", stop)

	// wait 5 second in background while server is trying to shut down
	ctxSvc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//try to shut down the server
	logger.Warn("shutting down http server")
	if err := server.Shutdown(ctxSvc); err != nil {
		logger.Error(err)
	}
}

package http

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type Users struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var menu = []Users{
	{
		ID:   "1",
		Name: "Freddie Mercury",
	},
	{
		ID:   "1",
		Name: "Brian May",
	},
	{
		ID:   "1",
		Name: "Roger Taylor",
	},
	{
		ID:   "1",
		Name: "John Deaconß",
	},
}

// Endpoints related functions
func listUsers(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(menu); err != nil {
		http.Error(w, "Unable to enconde data", 500)
	}
}

// Router related
func initRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/users", listUsers)

	return router
}

func InitHttpServer() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	router := initRouter()
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logrus.Info("HTTP server started. Logs for testing purposes only")

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		logrus.WithError(err).Panic("Server shutting down with error...")
	}
}

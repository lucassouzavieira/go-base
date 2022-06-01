package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
)

type item struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

var menu = []item{
	{ID: "1", Title: "Baby Ribs with Fries", Description: "Ribs with fries. Optional Barbecue sauce", Price: 2.35},
	{ID: "2", Title: "Chicken with Rice", Description: "Rice, fried chicken and Teriyaki sauce", Price: 2.35},
	{ID: "3", Title: "Vegetarian", Description: "Vegetarian meal", Price: 2.35},
	{ID: "4", Title: "Natural sandwich", Description: "Small sandwich with tomatoes, cheddar and ham.", Price: 2.35},
}

// Endpoints related functions
func listItems(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(menu); err != nil {
		http.Error(w, "Unable to enconde data", 500)
	}
}

// Router related
func initRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/hello", listItems)

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

	logger := log.New(os.Stderr, "HTTP Server: ", 0)
	logger.Println("HTTP server started. Logs for testing purposes only")

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		logger.Printf("Exit reason: %s \n", err)
	}
}

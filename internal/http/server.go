package http

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	json.NewEncoder(w).Encode(menu)
}

// Router related
func initRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/hello", listItems)

	return router
}

func InitHttpServer() {
	router := initRouter()
	logger := log.New(os.Stderr, "ONBOARDING-LVIEIRA: ", 0)
	logger.Output(2, "HTTP server started. Logs for testing purposes only")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		logger.Output(2, err.Error())
	}
}

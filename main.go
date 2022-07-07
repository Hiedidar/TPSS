package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func initRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/split-payments/compute", SplitHandler).Methods(http.MethodGet)
	return router
}

func main() {

	router := initRouter()

	log.Println("Starting Server On Port 8080")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		log.Panicln(err)
	}
}

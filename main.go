package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maha-1030/go-movie-ratings/handlers"
	"github.com/maha-1030/go-movie-ratings/repo"
)

const (
	POPULATE_DATA_FLAG_NAME        = "populateData"
	POPULATE_DATA_FLAG_DESCRIPTION = "send it as true if you want to populate fresh data in db"
)

func main() {
	populateDataFlag := flag.Bool(POPULATE_DATA_FLAG_NAME, false, POPULATE_DATA_FLAG_DESCRIPTION)

	flag.Parse()

	if populateDataFlag != nil && *populateDataFlag {
		repo.PopulateData()
	}

	router := mux.NewRouter().Headers("Content-Type", "application/json").PathPrefix("/api/v1").Subrouter()

	router.HandleFunc("/new-movie", handlers.NewMovieHandler).Methods(http.MethodPost)
	router.HandleFunc("/longest-duration-movies", handlers.GetLongestDurationMoviesHandler).Methods(http.MethodGet)

	http.ListenAndServe("localhost:8080", router)
}

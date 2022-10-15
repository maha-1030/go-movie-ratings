package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maha-1030/go-movie-ratings/repo"
)

const (
	RUNTIME_MINUTES_FIELD = "runtime_minutes"
	NO_OF_TOP_MOVIES      = 10
	DESCENDING_ORDER      = "DESC"
)

type NewMovieRequest struct {
	Movie  *repo.Movie
	Rating *repo.Rating
}

func NewMovieHandler(w http.ResponseWriter, r *http.Request) {
	newMovieRequest := NewMovieRequest{}

	if err := json.NewDecoder(r.Body).Decode(&newMovieRequest); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "invalid request body")

		return
	}

	if newMovieRequest.Movie == nil {
		respondWithError(w, http.StatusBadRequest, "movie details are not found")

		return
	}

	if newMovieRequest.Rating == nil {
		respondWithError(w, http.StatusBadRequest, "movie rating is not found")

		return
	}

	if err := newMovieRequest.Movie.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())

		return
	}

	if err := newMovieRequest.Rating.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())

		return
	}

	if newMovieRequest.Movie.Tconst != newMovieRequest.Rating.Tconst {
		respondWithError(w, http.StatusBadRequest, "Tconst must be same for both movie and rating of that movie")

		return
	}

	if m, err := (&repo.Movie{}).GetByTconst(newMovieRequest.Movie.Tconst); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	} else if m != nil {
		respondWithError(w, http.StatusBadRequest, "movie already exists with given tconst, use unique tconst")

		return
	}

	if r, err := (&repo.Rating{}).GetByTconst(newMovieRequest.Movie.Tconst); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	} else if r != nil {
		respondWithError(w, http.StatusBadRequest, "rating already exists with given tconst, use unique tconst")

		return
	}

	if err := newMovieRequest.Movie.Save(); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	if err := newMovieRequest.Rating.Save(); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "internal server error")

		return
	}

	respondWithJson(w, http.StatusOK, "success")
}

func GetLongestDurationMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movie := &repo.Movie{}

	longestDurationMovies, err := movie.GetTopMoviesByField(RUNTIME_MINUTES_FIELD, DESCENDING_ORDER, NO_OF_TOP_MOVIES)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")

		return
	}

	respondWithJson(w, http.StatusOK, longestDurationMovies)
}

func respondWithJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func respondWithError(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": errMsg})
}

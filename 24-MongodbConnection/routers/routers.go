package routers

import (
	controllers "github.com/nikhilrana/Golang-Tutorials/24-MongodbConnection/controllers"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMoviesService).Methods("GET")
	router.HandleFunc("/api/movie", controllers.AddMovieService).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkedMovieAsWatchedService).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteMovieService).Methods("DELETE")
	router.HandleFunc("/api/delete/all", controllers.DeleteAllMoviesService).Methods("DELETE")

	return router
}

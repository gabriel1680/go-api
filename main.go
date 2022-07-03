package main

import (
	"fmt"
	"log"
	"movies-api/adapter"
	"movies-api/repository"
	use_case "movies-api/use-case"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	repo := repository.NewMovieRepository()
	createUseCase := use_case.NewCreateMovieUseCase(repo)
	updateUseCase := use_case.NewUpdateMovieUseCase(repo)
	findByIdUseCase := use_case.NewFindByIdMovieUseCase(repo)
	listUseCase := use_case.NewListMovieUseCase(repo)
	deleteUseCase := use_case.NewDeleteMovieUseCase(repo)

	router := mux.NewRouter()
	router.HandleFunc("/{id}", adapter.FindRouteAdapter(findByIdUseCase)).Methods("GET")
	router.HandleFunc("/", adapter.ListRouteAdapter(listUseCase)).Methods("GET")
	router.HandleFunc("/", adapter.CreateRouteAdapter(createUseCase)).Methods("POST")
	router.HandleFunc("/", adapter.UpdateRouteAdapter(updateUseCase)).Methods("PUT")
	router.HandleFunc("/{id}", adapter.DeleteRouteAdapter(deleteUseCase)).Methods("DELETE")

	fmt.Println("Start server on port :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

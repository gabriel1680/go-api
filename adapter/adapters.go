package adapter

import (
	"encoding/json"
	"net/http"
	"strconv"

	"movies-api/entity"
	use_case "movies-api/use-case"

	"github.com/gorilla/mux"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type route func(http.ResponseWriter, *http.Request)

func FindRouteAdapter(uc use_case.FindByIdMovieUseCase) route {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		response, err := uc.Execute(id)

		rw.Header().Set("Content Type", "application/json")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
		} else {
			rw.WriteHeader(http.StatusOK)
			msg, _ := json.Marshal(response)
			rw.Write(msg)
		}

	}
}

func ListRouteAdapter(uc use_case.ListMovieUseCase) route {
	return func(rw http.ResponseWriter, r *http.Request) {
		response, err := uc.Execute()

		rw.Header().Set("Content Type", "application/json")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
		} else {
			rw.WriteHeader(http.StatusOK)
			msg, _ := json.Marshal(response)
			rw.Write(msg)
		}

	}
}

func DeleteRouteAdapter(cb use_case.DeleteMovieUseCase) route {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		err := cb.Execute(id)

		rw.Header().Set("Content Type", "application/json")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
		} else {
			rw.WriteHeader(http.StatusNoContent)
			msg, _ := json.Marshal(&Response{Status: "success", Message: "deleted"})
			rw.Write(msg)
		}

	}
}

func CreateRouteAdapter(cb use_case.CreateMovieUseCase) route {
	return func(rw http.ResponseWriter, r *http.Request) {
		var movie entity.Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)
		err := cb.Execute(&movie)

		rw.Header().Set("Content Type", "application/json")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
		} else {
			rw.WriteHeader(http.StatusCreated)
			msg, _ := json.Marshal(&Response{Status: "success", Message: "created"})
			rw.Write(msg)
		}
	}
}

func UpdateRouteAdapter(cb use_case.UpdateMovieUseCase) route {
	return func(rw http.ResponseWriter, r *http.Request) {
		var movie entity.Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)
		err := cb.Execute(&movie)

		rw.Header().Set("Content Type", "application/json")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
		} else {
			rw.WriteHeader(http.StatusOK)
			msg, _ := json.Marshal(&Response{Status: "success", Message: "updated"})
			rw.Write(msg)
		}
	}
}

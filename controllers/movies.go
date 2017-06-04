package controllers

import (
	"encoding/json"
	//"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"moviesapp/models"
	"moviesapp/token"
	"net/http"
)

type errOut struct {
	Error string `json:error`
}

func GetMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(models.GetMovies()); err != nil {
		log.Panic("Error EncodingJson in ControllersGetMovies", err)

	}
}

func NewMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	status := token.TokenHandler(w, r)
	if status == true {
		id, err := models.NewMovie(r)
		if err == nil && id != "" {
			if err := json.NewEncoder(w).Encode(models.GetMovieById(id)); err != nil {
				log.Panic("Error EncodingJson in ControllersNewMovie", err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				log.Panic("Error EncodingJson in ControllersNewMovie", err)
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func GetMovieById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(models.GetMovieById(p.ByName("id"))); err != nil {
		log.Panic("Controller GetMovieById json err: ", err)
	}
}

func UpdateMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	status := token.TokenHandler(w, r)
	if status == true {
		id, err := models.UpdateMovie(r, p.ByName("id"))
		if err == nil {
			if err := json.NewEncoder(w).Encode(id); err != nil {
				log.Panic("Controller UpdateMovieById json err: ", err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				log.Panic("Controller UpdateMovieById json err: ", err)
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	status := token.TokenHandler(w, r)
	if status == true {
		if err := json.NewEncoder(w).Encode(models.DeleteMovie(p.ByName("id"))); err != nil {
			log.Panic("Controller UpdateMovieById json err: ", err)
		}
	}
}

package controllers

import (
	"encoding/json"
	//	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"moviesapp/models"
	"net/http"
)

func GetMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(models.GetMovies()); err != nil {
		log.Panic("Error EncodingJson in ControllersGetMovies", err)
	}

}

func NewMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(models.NewMovie(r)); err != nil {
		log.Panic("Error EncodingJson in ControllersNewMovie", err)
	}
}

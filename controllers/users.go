package controllers

import (
	"encoding/json"
	//"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"moviesapp/models"
	"net/http"
	//"reflect"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res, err := models.Login(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Panic("Controller User: Login json err: ", err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Panic("Controller User: Register Json err: ", err)
		}
	}
	//fmt.Println(reflect.TypeOf(r.Body))
}

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res, err := models.Register(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Panic("Controller User: Register Json err: ", err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Panic("Controller User: Register Json err: ", err)
		}
	}
}

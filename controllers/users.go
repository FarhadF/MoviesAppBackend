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
	if err := json.NewEncoder(w).Encode(models.Login(r)); err != nil {
		log.Panic("Controller User: Login json err: ", err)
	}
	//fmt.Println(reflect.TypeOf(r.Body))
}

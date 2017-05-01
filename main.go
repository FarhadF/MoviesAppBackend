package main

import (
	//	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"moviesapp/controllers"
	"moviesapp/models"
	"net/http"
)

func main() {
	models.InitDB()
	router := httprouter.New()
	router.GET("/movies", controllers.GetMovies)
	router.POST("/movie/new", controllers.NewMovie)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Panic("ListenAndServeErr: ", err)
	}

}

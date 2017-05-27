package main

import (
	//	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"moviesapp/controllers"
	"moviesapp/logging"
	"moviesapp/models"

	"net/http"
	"os"
)

func main() {
	models.InitDB()
	router := httprouter.New()
	router.GET("/movies", controllers.GetMovies)
	router.POST("/movie/", controllers.NewMovie)
	router.GET("/movie/:id", controllers.GetMovieById)
	router.POST("/movie/:id/edit", controllers.UpdateMovie)
	router.DELETE("/movie/:id", controllers.DeleteMovie)
	router.POST("/login", controllers.Login)
	//handler := cors.Default().Handler(router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
	})
	handler := c.Handler(router)
	loggingHandler := logging.NewApacheLoggingHandler(handler, os.Stderr)
	server := &http.Server{
		Addr:    ":8080",
		Handler: loggingHandler,
	}
	//err := http.ListenAndServe(":8080", router)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic("ListenAndServeErr: ", err)
	}

}

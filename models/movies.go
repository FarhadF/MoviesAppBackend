package models

import (
	"database/sql"
	//"fmt"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"time"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("mysql", "farhad:Aa111111@tcp(localhost:3306)/go?charset=utf8")
	if err != nil {
		log.Panic("InitDB: ", err)
	}
	if err = Db.Ping(); err != nil {
		log.Panic("Db.Ping Failed: ", err)
	}

}

type Movie struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Year      string `json:"year"`
	Director  string `json:"director"`
	Timestamp string `json:"timestamp"`
}

func GetMovies() []*Movie {
	rows, err := Db.Query("select * from movies")
	if err != nil {
		log.Panic("GetMovies Select Failed: ", err)
	}
	Movies := make([]*Movie, 0)
	for rows.Next() {
		movie := new(Movie)
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Director, &movie.Timestamp)
		if err != nil {
			log.Panic("Movie Struct failed: ", err)
		}

		Movies = append(Movies, movie)
	}
	return Movies

}

type ErrOut struct {
	Error string `json:"error"`
}

func NewMovie(r *http.Request) interface{} {
	t := time.Now()
	movie := new(Movie)
	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil {
		log.Panic("JsonDecode failed in ModelsNewMovie: ", err)
	}

	rows, err := Db.Query("select * from movies where name='" + movie.Name + "'")
	if err != nil {
		log.Panic("Model New Movie Select failed: ", err)
	}
	if !rows.Next() {
		stmt, err := Db.Prepare("insert into movies set name=?, year=?, director=?, timestamp=?")
		res, err := stmt.Exec(movie.Name, movie.Year, movie.Director, t.Format("2006-01-02 15:04:05"))
		id, err := res.LastInsertId()
		if err != nil {
			log.Panic("Model NewMovie insert failed: ", err)
		}
		return strconv.FormatInt(id, 10)
	} else {
		errout := new(ErrOut)
		errout.Error = "Duplicate Name"
		return errout
	}
}

package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	//	"time"
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
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Year     string `json:"year"`
	Director string `json:"director"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

func GetMovies() []*Movie {
	rows, err := Db.Query("select * from movies")
	if err != nil {
		log.Panic("GetMovies Select Failed: ", err)
	}
	Movies := make([]*Movie, 0)
	for rows.Next() {
		movie := new(Movie)
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Director, &movie.Created, &movie.Updated)
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

func NewMovie(r *http.Request) (string, *ErrOut) {
	//t := time.Now()
	movie := new(Movie)
	err := json.NewDecoder(r.Body).Decode(&movie)
	if movie.Name != "" && movie.Director != "" && movie.Year != "" {
		if err != nil {
			log.Panic("JsonDecode failed in ModelsNewMovie: ", err)
		}
		rows, err := Db.Query("select * from movies where name='" + movie.Name + "'")
		if err != nil {
			log.Panic("Model New Movie Select failed: ", err)
		}
		if !rows.Next() {
			stmt, err := Db.Prepare("insert into movies set name=?, year=?, director=?")
			res, err := stmt.Exec(movie.Name, movie.Year, movie.Director) //, t.Format("2006-01-02 15:04:05"))
			id, err := res.LastInsertId()
			if err != nil {
				log.Panic("Model NewMovie insert failed: ", err)
			}
			return strconv.FormatInt(id, 10), nil
		} else {
			errout := new(ErrOut)
			errout.Error = "duplicate name"
			return "", errout
		}
	} else {
		errout := new(ErrOut)
		errout.Error = "fill the required fields"
		return "", errout
	}
}

func GetMovieById(id string) interface{} {
	movie := new(Movie)
	err := Db.QueryRow("select * from movies where id='"+id+"'").Scan(&movie.Id, &movie.Name, &movie.Year, &movie.Director, &movie.Created, &movie.Updated)
	if err != nil && err.Error() != "sql: no rows in result set" {
		log.Panic("Model GetMovieById Select failed: ", err)
	}
	if err != nil && err.Error() == "sql: no rows in result set" {
		fmt.Println("No Data found in Database!")
		errout := new(ErrOut)
		errout.Error = "Requested ID not found"
		return errout
	} else {
		return movie
	}

}

func UpdateMovie(r *http.Request, id string) interface{} {
	movie := new(Movie)
	//t := time.Now()
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Panic("JsonDecode failed in ModelsEditMovie: ", err)
	}
	stmt, err := Db.Prepare("update movies set name=?, year=?, director=? where id=?")
	res, err := stmt.Exec(movie.Name, movie.Year, movie.Director, id) //t.Format("2006-01-02 15:04:05"), id)
	affected, err := res.RowsAffected()
	if err != nil {
		log.Panic("DB Error in ModelsEditMovie: ", err)
	}
	return strconv.FormatInt(affected, 10)

}

func DeleteMovie(id string) interface{} {
	stmt, err := Db.Prepare("delete from movies where id=?")
	res, err := stmt.Exec(id)
	affected, err := res.RowsAffected()
	if err != nil {
		log.Panic("DB Error in ModelsDeleteMovie: ", err)
	}
	return strconv.FormatInt(affected, 10)
}

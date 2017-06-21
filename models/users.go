package models

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	//"strconv"
	//"time"
	"moviesapp/token"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorOut struct {
	Error string `json:"error"`
}

type TokenOut struct {
	Token string `json:"token"`
}

func Login(cred *http.Request) interface{} {

	var usercred UserCredentials
	var user User
	err := json.NewDecoder(cred.Body).Decode(&usercred)
	if err != nil {
		log.Println("Models Users Login json failed: ", err)
		return `{"error": "Server Error"}`
	}
	fmt.Println(usercred)
	err = Db.QueryRow("select * from users where email='"+usercred.Email+"'").Scan(&user.Id, &user.Email, &user.Name, &user.Lastname, &user.Password, &user.Role, &user.Created, &user.Updated)
	if err != nil && err.Error() != "sql: no rows in result set" {
		log.Println("Models users login query Failed: ", err)
	}
	if err != nil && err.Error() == "sql: no rows in result set" {
		log.Println("User Not Found")
		errOut := new(ErrorOut)
		errOut.Error = "Invalid Username or Password"
		return errOut
	}
	if err == nil && user.Password != usercred.Password {
		fmt.Println(user)
		log.Println("Invalid Password:", usercred.Password, user.Password)
		errOut := new(ErrorOut)
		errOut.Error = "Invalid Username or Password"
		return errOut

	} else {
		//do token stuff
		generatedToken, err := token.GenerateToken(user.Role)
		if err != nil {
			log.Println("Token creation failed: ", err)
			errOut := new(ErrorOut)
			errOut.Error = "Invalid Username or Password"
			return errOut
		}
		jsonToken := new(TokenOut)
		jsonToken.Token = generatedToken
		return jsonToken
	}

}

func Register(r *http.Request) interface{} {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		errOut := new(ErrorOut)
		errOut.Error = "Server Error"
	}
	if user.Email == "" {
		log.Println("Please enter your Email.")
		errOut := new(ErrorOut)
		errOut.Error = "Please enter your Email."
		return errOut
	} else {
		if user.Lastname == "" {
			log.Println("Please enter your Lastname.")
			errOut := new(ErrorOut)
			errOut.Error = "Please enter your Lastname."
			return errOut
		}
		if user.Name == "" {
			log.Println("Please enter your Name.")
			errOut := new(ErrorOut)
			errOut.Error = "Please enter your Name."
			return errOut
		}
		if user.Password == "" {
			log.Println("Please enter your Password.")
			errOut := new(ErrorOut)
			errOut.Error = "Please enter your Password."
			return errOut
		}
		var exists bool
		stmt, err := Db.Prepare("select exists (select * from users where email=?)")
		err = stmt.QueryRow(user.Email).Scan(&exists)
		if err != nil {
			log.Println("Register Query failed: ", err)
		}
		if exists == true {
			log.Println("Email already exists")
			errOut := new(ErrorOut)
			errOut.Error = "Email already registered"
			return errOut
		} else {
			//Register User
			stmt, err := Db.Prepare("insert into users set email=?,name=?,lastname=?,password=?,role=?")
			if err != nil {
				log.Println("Register Query failed: ", err)
			}
			res, err := stmt.Exec(user.Email, user.Name, user.Lastname, user.Password, "user")
			if err != nil {
				log.Println("Register Query failed: ", err)
			}
			id, err := res.LastInsertId()
			return id
		}
	}
}

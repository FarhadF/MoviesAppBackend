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
)

type User struct {
	Id       int    `json:id`
	Name     string `json:name`
	Lastname string `json:lastname`
	Email    string `json:email`
	Password string `json:password`
	Role     string `json:role`
	Created  string `json:created`
	Updated  string `json:Updated`
}

type UserCredentials struct {
	Email    string `json:email`
	Password string `json:password`
}

type ErrorOut struct {
	Error string `json:error`
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
	err = Db.QueryRow("select * from users where email='"+usercred.Email+"'").Scan(&user.Id, &user.Password, &user.Name, &user.Lastname, &user.Email, &user.Created, &user.Updated, &user.Role)
	if err != nil && err.Error() != "sql: no rows in result set" {
		log.Println("Models users login query Failed: ", err)
	}
	if err != nil && err.Error() == "sql: no rows in result set" {
		log.Println("User Not Found")
		errOut := new(ErrorOut)
		errOut.Error = "Invalid Username or Password"
		return errOut
	}
	if user.Password != usercred.Password {
		fmt.Println(user)
		log.Println("Invalid Password:", usercred.Password, user.Password)
		errOut := new(ErrorOut)
		errOut.Error = "Invalid Username or Password"
		return errOut

	}
	fmt.Println(user)
	return user
}

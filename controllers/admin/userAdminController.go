package controllers

import (
	"backend_api/database"
	"backend_api/models"
	"backend_api/repositores"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// UserAdminController is the controller for the user admin routes
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositores.NewRepositoryUsers(db)
	userID, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}
	w.Write([]byte(fmt.Sprintf("User created with ID: %d", userID)))
}

// UserAdminController is the controller for the user admin routes
func GetUsersAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}

// UserAdminController is the controller for the user admin routes
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

// UserAdminController is the controller for the user admin routes
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

// UserAdminController is the controller for the user admin routes
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}

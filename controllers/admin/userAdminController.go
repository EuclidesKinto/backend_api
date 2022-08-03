package controllers

import (
	"backend_api/database"
	messages "backend_api/messages"
	"backend_api/models"
	"backend_api/repositores"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// UserAdminController is the controller for the user admin routes
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		messages.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare(); erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositores.NewRepositoryUsers(db)
	userID, erro := repository.Create(user)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	messages.JSON(w, http.StatusCreated, userID)
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

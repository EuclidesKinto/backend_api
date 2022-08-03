package controllers

import (
	"backend_api/config"
	"backend_api/database"
	"backend_api/messages"
	"backend_api/models"
	"backend_api/repositores"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// RouteLogin is the route for the login
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, erro := database.Connect()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositores.NewRepositoryUsers(db)
	userSaved, erro := repository.GetByEmail(user.Email)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = config.VerifyPassword(userSaved.Password, user.Password); erro != nil {
		messages.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	w.Write([]byte("Você está logado"))
}

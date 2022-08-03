package controllers

import "net/http"

// UserAdminController is the controller for the user admin routes
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateUser"))
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

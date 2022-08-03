package routes

import "github.com/gorilla/mux"

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return ConfigRouter(r)
}

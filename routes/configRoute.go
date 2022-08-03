package routes

import (
	"backend_api/routes/api/admin"
	"backend_api/routes/api/auth"

	"github.com/gorilla/mux"
)

// Config coloca todas as routes dentro de um router
func ConfigRouter(r *mux.Router) *mux.Router {
	routes := admin.RoutesUsers
	routes = append(routes, auth.RouteLogin)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.HandlerFunc).Methods(route.Method)
	}

	return r
}

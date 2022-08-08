package routes

import (
	"backend_api/middleware"
	"backend_api/routes/api/admin"
	"backend_api/routes/api/auth"

	"github.com/gorilla/mux"
)

// Config coloca todas as routes dentro de um router
func ConfigRouter(r *mux.Router) *mux.Router {
	routes := admin.RoutesUsers
	routes = append(routes, auth.RouteLogin)

	for _, route := range routes {
		if route.RequeriAuth {
			r.HandleFunc(route.URI,
				middleware.Logger(middleware.Authentication(route.HandlerFunc))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middleware.Logger(route.HandlerFunc)).Methods(route.Method)

		}
	}

	return r
}

package admin

import (
	controllers "backend_api/controllers/admin"
	"backend_api/routes/api"
	"net/http"
)

var url = "/api/admin/"

var RoutesUsers = []api.Route{
	{
		URI:         url + "users",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateUser,
		RequeriAuth: false,
	},
	{
		URI:         url + "users",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUsersAll,
		RequeriAuth: false,
	},
	{
		URI:         url + "users/{id}",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUser,
		RequeriAuth: false,
	},
	{
		URI:         url + "users/{id}",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateUser,
		RequeriAuth: false,
	},
	{
		URI:         url + "users/{id}",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteUser,
		RequeriAuth: false,
	},
}

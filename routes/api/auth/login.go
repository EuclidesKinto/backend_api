package auth

import (
	controllers "backend_api/controllers/auth"
	"backend_api/routes/api"
	"net/http"
)

var url = "/api/auth/"

var RouteLogin = api.Route{

	URI:         url + "login",
	Method:      http.MethodPost,
	HandlerFunc: controllers.Login,
	RequeriAuth: false,
}

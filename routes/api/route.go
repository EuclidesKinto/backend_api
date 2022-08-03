package api

import (
	"net/http"
)

// route is a struct. All routes are stored in this struct.
type Route struct {
	URI         string
	Method      string
	HandlerFunc func(http.ResponseWriter, *http.Request)
	RequeriAuth bool
}

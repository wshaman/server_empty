package server

import (
	"github.com/wshaman/server_empty/api/v1/handlers"
	"net/http"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}

// Routes - A collection of Routes
type Routes []Route

// APIRoutes - Routes for ShitHappens API
var apiRoutes = Routes{
	Route{
		http.MethodPost,
		"/role",
		handlers.RoleAdd,
		false,
	},
	Route{
		http.MethodGet,
		"/role",
		handlers.RoleList,
		true,
	},
}

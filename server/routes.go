package server

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var apiRoutes = Routes{
	// Lists
	Route{
		"ListNames",
		"GET",
		"/names",
		ListNames,
	},
	Route{
		"ListIds",
		"GET",
		"/ids",
		ListIds,
	},
	Route{
		"ListDomains",
		"GET",
		"/domains",
		ListDomains,
	},
	Route{
		"ListExpiredNames",
		"GET",
		"/expired",
		ListExpiredNames,
	},
	// Show identifiers
	Route{
		"ShowId",
		"GET",
		"/ids/{identifier}",
		ShowId,
	},
	Route{
		"ShowDomain",
		"GET",
		"/domains/{identifier}",
		ShowDomain,
	},
	// Search
	Route{
		"FilterNames",
		"POST",
		"/search",
		FilterNames,
	},
	// Balance
	Route{
		"GetBalance",
		"GET",
		"/balance",
		GetBalance,
	},
}

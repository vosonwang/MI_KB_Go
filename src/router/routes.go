package router

import (
	"github.com/julienschmidt/httprouter"
	"handlers"
	"log"
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Catalog,
	},
}

func init()  {
	router := httprouter.New()

	for _, route := range routes {

		router.Handle(route.Method, route.Pattern, route.Handle)
	}

	log.Fatal(http.ListenAndServe(":8081", router))

}
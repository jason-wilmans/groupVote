package routeRegistry

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/muroc/prego"
	"log"
)

var	router = mux.NewRouter().StrictSlash(true)
var	routeMap = make(map[string]bool)

func Register(pattern string, method string, handler http.HandlerFunc) {
	precond.NotNil(pattern, "")
	precond.NotNil(method, "")
	precond.NotNil(handler, "")
	precond.False(HasRoute(pattern), "")

	router.HandleFunc(pattern, handler)
	routeMap[pattern] = true
	log.Println("Registered Route: ", pattern)
}

func HasRoute(name string) bool {
	_, exists := routeMap[name]
	return exists
}

func StartHosting()  {
	log.Fatal(http.ListenAndServe(":8000", router))
}
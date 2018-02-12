package routeRegistry

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/muroc/prego"
	"log"
	"strconv"
)

const (
	GET = iota
	POST = iota
	PUT = iota
	DELETE = iota
)

var	router = mux.NewRouter().StrictSlash(true)
var	routeMap = make(map[string]bool)

func Register(path string, method int, handler http.HandlerFunc) {
	precond.NotNil(path, "")
	precond.InRange(float64(method), GET, DELETE, "Invalid http method value.")
	precond.NotNil(handler, "")
	precond.False(HasRoute(path, method), "")

	router.HandleFunc(path, handler)
	routeMap[path + strconv.Itoa(method)] = true
	log.Println("Registered Route: ", path, " (" , method, ")")
}

func HasRoute(path string, method int) bool {
	precond.NotNil(path, "")
	precond.InRange(float64(method), GET, DELETE, "Invalid http method value.")

	_, exists := routeMap[path + strconv.Itoa(method)]
	return exists
}

func StartHosting()  {
	log.Fatal(http.ListenAndServe(":8080", router))
}
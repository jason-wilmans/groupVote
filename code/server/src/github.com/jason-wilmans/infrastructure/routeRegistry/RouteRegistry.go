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

	var methodName = convertMethod(method)
	router.HandleFunc(path, handler).Methods(methodName, "OPTIONS")
	routeMap[path + strconv.Itoa(method)] = true
	log.Println("Registered Route: ", path, "(",methodName, ")")
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

func convertMethod(method int) string {
	precond.InRange(float64(method), GET, DELETE, "Invalid http method value.")

	switch method {
		case GET: return "GET"
		case POST: return "POST"
		case PUT: return "PUT"
		case DELETE: return "DELETE"
		default: return ""
	}
}
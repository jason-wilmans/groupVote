package routeRegistry

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/muroc/prego"
	"log"
	"bytes"
)

const (
	OPTIONS = iota
	GET = iota
	POST = iota
	PUT = iota
	DELETE = iota
)

var	router = mux.NewRouter().StrictSlash(true)
var	routeMap = make(map[string][]int)

func Register(path string, method int, handler http.HandlerFunc) {
	precond.NotNil(path, "")
	precond.InRange(float64(method), GET, DELETE, "Invalid http methodName.")
	precond.NotNil(handler, "")
	precond.False(HasRoute(path, method), "")

	var methodName = methodName(method)
	methods, exists := routeMap[path]
	if !exists {
		router.HandleFunc(path, enableCORS(path, handler)).Methods("OPTIONS")
		methods = make([]int, 0)
		routeMap[path] = methods
	}

	if alreadyContains(methods, method) {
		panic("The route' " + path + "'[" + methodName + "] is already registered.")
	}

	router.HandleFunc(path, enableCORS(path, handler)).Methods(methodName)
	routeMap[path] = append(methods, method)
	log.Println("Registered Route: ", path, "[",methodName,"]")
}

func enableCORS(path string, handler http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", listMethods(path))

		if r.Method == methodName(OPTIONS) {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler(w, r)
	}
}

func listMethods(path string) string {
	methods, exists := routeMap[path]
	if !exists {
		return ""
	}

	size := len(methods)
	var buffer bytes.Buffer
	for i, method := range methods {
		buffer.WriteString(methodName(method))
		if i < size-1 {
			buffer.WriteString(",")
		}
	}

	return buffer.String()
}

func HasRoute(path string, method int) bool {
	precond.NotNil(path, "")
	precond.InRange(float64(method), GET, DELETE, "Invalid http methodName value.")

	methods, exists := routeMap[path]
	return exists && alreadyContains(methods, method)
}

func StartHosting()  {
	log.Fatal(http.ListenAndServe(":8080", router))
}

func methodName(method int) string {
	precond.InRange(float64(method), OPTIONS, DELETE, "Invalid http methodName value.")

	switch method {
		case OPTIONS: return "OPTIONS"
		case GET: return "GET"
		case POST: return "POST"
		case PUT: return "PUT"
		case DELETE: return "DELETE"
		default: return ""
	}
}

func alreadyContains(methods []int, method int) bool {
	for _, value := range methods {
		if value == method {
			return true
		}
	}

	return false
}
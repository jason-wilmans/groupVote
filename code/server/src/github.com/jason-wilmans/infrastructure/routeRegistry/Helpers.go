package routeRegistry

import (
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
	"fmt"
	)

func ReadIdParam(paramKey string, writer http.ResponseWriter, request *http.Request) (uuid.UUID, error) {
	idString := mux.Vars(request)[paramKey]
	id, err := uuid.FromString(idString)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, err)
		return uuid.Nil, err
	}
	return id, error(nil)
}
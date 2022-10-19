package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mynoveloper/logger"
)

// var emptyArray = [...]string{}
// var emptyObject = make(map[string]string)

type httpResponse struct {
	logger logger.ILogger
}

type HttpResponse interface {
	Respond(w http.ResponseWriter, value interface{}, statusCode int)
	RespondError(w http.ResponseWriter, err interface{})
}

func NewResponse(logger logger.ILogger) HttpResponse {
	return &httpResponse{
		logger: logger,
	}
}

func (r *httpResponse) Respond(w http.ResponseWriter, value interface{}, statuscode int) {
	b, err := json.Marshal(value)

	if err != nil {
		r.RespondError(w, fmt.Errorf("No se puede obtener la respuesta: %v", err))
		return
	}

	w.WriteHeader(statuscode)
	w.Write(b)
}

func (r *httpResponse) RespondError(w http.ResponseWriter, err interface{}) {
	r.logger.Error(err)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

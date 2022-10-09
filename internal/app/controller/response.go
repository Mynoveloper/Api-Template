package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mynoveloper/logger"
)

// var emptyArray = [...]string{}
// var emptyObject = make(map[string]string)

type httpResponse struct{}
type IHttpResponse interface {
	Respond(w http.ResponseWriter, value interface{}, statusCode int)
	RespondError(w http.ResponseWriter, err interface{})
}

var _logger logger.ILogger

func NewResponse(logger logger.ILogger) IHttpResponse {
	_logger = logger
	return &httpResponse{}
}

func (c *httpResponse) Respond(w http.ResponseWriter, value interface{}, statuscode int) {
	b, err := json.Marshal(value)

	if err != nil {
		c.RespondError(w, fmt.Errorf("No se puede obtener la respuesta: %v", err))
		return
	}

	w.WriteHeader(statuscode)
	w.Write(b)
}

func (c *httpResponse) RespondError(w http.ResponseWriter, err interface{}) {
	_logger.Error(err)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

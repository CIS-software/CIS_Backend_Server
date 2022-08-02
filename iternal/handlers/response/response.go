package response

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//Error fills the response body with an error and redirects the data to the responder
func Error(w http.ResponseWriter, code int, err error) {
	log.Info(err)
	if code == http.StatusUnprocessableEntity {
		err = errors.New(
			"the data type and syntax is correct, but for some reason the server cannot process the data")
	}
	Respond(w, code, map[string]string{"error": err.Error()})
}

//Respond returns a response to a client request in the form of status code and response body
func Respond(w http.ResponseWriter, code int, data interface{}) {
	log.Info("Status: ", code)
	log.Info(data)
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

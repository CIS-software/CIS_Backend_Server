package response

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//Error fills the response body with an error and redirects the data to the responder
func Error(w http.ResponseWriter, code int, err error) {
	log.Info("Err: ", err)

	if code == http.StatusUnprocessableEntity {
		err = errors.New("data type and structure are correct but can't process data")
	}

	Respond(w, code, map[string]string{"error": err.Error()})
}

//Respond returns a response to a client request in the form of status code and response body
func Respond(w http.ResponseWriter, code int, data interface{}) {
	//log status code and response body content
	log.Info("Status: ", code)
	log.Info("Response body: ", data)

	//return code status to client
	w.WriteHeader(code)

	//body return to client if not nil
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

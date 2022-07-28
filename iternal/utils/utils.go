package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	logrus.Error(err)
	Respond(w, r, code, map[string]string{"error": err.Error()})
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	logrus.Info("Status: ", code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

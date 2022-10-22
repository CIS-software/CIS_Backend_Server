package apiserver

import (
	"CIS_Backend_Server/iternal/handlers/response"
	"CIS_Backend_Server/iternal/model"
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
)

// middleware errors
var (
	ErrMissingAuthToken = errors.New("missing auth token")
	ErrWrongFormat      = errors.New("invalid auth token, does not match the format: Bearer {token-body}")
	ErrWrongToken       = errors.New("wrong token")
)

// LogRequest getting basic information about the request in the form of logs, such as method, url, header
func (s *Server) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		s.log.Infof("Request: Method: %s, URL: %s, Header: %s", r.Method, r.URL.Path, r.Header)
		next.ServeHTTP(w, r)
		s.log.Infof("Request time: %s", time.Since(start))
	})
}

// JwtAuthentication authorization with jwt tokens
func (s *Server) JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//array of urls that do not require authorization
		notAuth := []string{"/login", "/create-user", "/update-tokens"}

		//getting url address from client
		requestPath := r.URL.Path

		//checking the need for authorization of the received url
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		//getting token from header
		tokenHeader := r.Header.Get("Authorization")

		//check for the presence of a token
		if tokenHeader == "" {
			response.Error(w, http.StatusUnauthorized, ErrMissingAuthToken)
			return
		}

		//authorization format check
		indented := strings.Split(tokenHeader, " ")
		if len(indented) != 2 {
			response.Error(w, http.StatusUnauthorized, ErrWrongFormat)
			return
		}

		//token validation
		tokenPart := indented[1]
		t := &model.Tokens{}
		_, err := jwt.ParseWithClaims(tokenPart, t, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.secretKey), nil
		})

		//token is not valid
		if err != nil {
			response.Error(w, http.StatusUnauthorized, ErrWrongToken)
			return
		}

		ctx := context.WithValue(r.Context(), "user", t.TokenId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

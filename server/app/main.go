package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/sawadashota/hcmcafe/server/handler"
	"google.golang.org/appengine"
)

func init() {
	s := rpc.NewServer()

	s.RegisterCodec(json2.NewCodec(), "application/json")

	// CORS config
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	cors := handlers.CORS(headersOk, originOk, methodsOk, handlers.MaxAge(3600))

	// Register methods
	s.RegisterService(new(handler.Admin), "Admin")
	s.RegisterService(new(handler.Cafe), "Cafe")
	s.RegisterService(new(handler.HealthCheck), "HealthCheck")

	http.Handle("/rpc", cors(s))
}

func main() {
	appengine.Main()
}

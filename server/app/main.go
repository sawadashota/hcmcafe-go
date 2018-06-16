package main

import (
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/sawadashota/hcmcafe/server/handler"
	"google.golang.org/appengine"
)

func init() {
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")

	s.RegisterService(new(handler.Cafe), "Cafe")
	s.RegisterService(new(handler.HealthCheck), "HealthCheck")

	http.Handle("/rpc", s)
}

func main() {
	appengine.Main()
}

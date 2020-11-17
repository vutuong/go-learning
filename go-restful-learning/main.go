package main

import (
	"log"
	"net/http"
	"go-restful-learning/bookservice"
	"github.com/emicklei/go-restful"
)

func main() {
	// A Container holds a collection of WebServices, Filters and a http.ServeMux for multiplexing http requests
	// Option 1 - Using default container
	// restful.Add(ws)
	// log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil))

	// Option 2 - Define your own container
	// WebServices must be added to a container (see below) in order to handler Http requests from a server
	container := restful.NewContainer()
	ws := bookservice.NewAPIServer() 
	container.Add(ws)
	log.Fatal(http.ListenAndServe("0.0.0.0:8888", container))
}

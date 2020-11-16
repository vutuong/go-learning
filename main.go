package main

import (
	"log"
	"net/http"
	"go-restful-learning/bookservice"
	"github.com/emicklei/go-restful"
)

func main() {
	ws := bookservice.NewAPIServer()
	restful.Add(ws)
	log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil))
}

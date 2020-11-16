package bookservice

import (
	"bytes"
	"net/http"

	"github.com/emicklei/go-restful"
)

//Book defines book object
type Book struct {
	ID     string
	Title  string
	Author string
}

func init() {
	restful.MarshalIndent = func(v interface{}, prefix, indent string) ([]byte, error) {
		var buf bytes.Buffer
		encoder := restful.NewEncoder(&buf)
		encoder.SetIndent(prefix, indent)
		if err := encoder.Encode(v); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
}

// NewAPIServer Function New create a new API specification of the Webservice
// The API is a collection of Route objects which specify:
// how an incoming Http Request is mapped to a function
func NewAPIServer() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/books").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)
	service.Route(service.GET("/").To(mainPage))
	service.Route(service.GET("/{book-id}").To(findBook))
	service.Route(service.POST("").To(updateBook))
	service.Route(service.PUT("/{book-id}").To(createBook))
	service.Route(service.DELETE("/{book-id}").To(removeBook))

	return service
}

// The signature for functions that are used in Route is containing a restful Request.Response pair
// Request is a wrapper for a http Request that provides convenience methods
// Response is a wrapper on actual htpp ResponseWriter
// There are common restful functions such as WriteEntiy that will inspect the Accept header of
// the request to determine the Content-Type of the response and how to Marshal the object
// *** Cac functions duoc dung trong Route co dau vao chung la mot cap restful Request-Reponse
// *** WriteEntity la mot function restful pho bien duoc dung de kiem tra Accept header cua mot request
// de xac dinh Content-Type cua reponse cung nhu cach Marchal object

func mainPage(request *restful.Request, response *restful.Response) {
	hello := "Hello world!"
	response.WriteEntity(hello)
}

func findBook(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("book-id")
	// here you would fetch book from some persistence system
	book := &Book{ID: id, Title: "The Alchemist", Author: "Paulo Coelho"}
	response.WriteEntity(book)
}

func updateBook(request *restful.Request, response *restful.Response) {
	book := new(Book)
	err := request.ReadEntity(&book)
	// here update the book with some persistence system
	if err == nil {
		response.WriteEntity(book)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func createBook(request *restful.Request, response *restful.Response) {
	book := Book{ID: request.PathParameter("book-id")}
	err := request.ReadEntity(&book)
	// here you would create the book with some persistence system
	if err == nil {
		response.WriteEntity(book)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func removeBook(request *restful.Request, response *restful.Response) {
	// here you would delete the book from some persistence system
}

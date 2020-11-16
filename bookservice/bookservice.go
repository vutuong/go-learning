package bookservice

fmt (
	"github.com/emicklei/go-restful"

)

type book struct {
	ID int
	Title string
	Author string
}

// Function New create a new API specification of the Webservice
// The API is a collection of Route objects which specify:
// how an incoming Http Request is mapped to a function
func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/books").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)
		
	service.Route(service.GET("/{book-id}").To(FindBook))
	service.Route(service.POST("").To(UpdateBook))
	service.Route(service.PUT("/{book-id}").To(CreateBook))
	service.Route(service.DELETE("/{book-id}").To(RemoveBook))
		
	return service
}
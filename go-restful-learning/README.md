# The-first-go-restful-working-example
Creating restful API with golang

1. Webservice and Routes.
- Một Webservice bao gồm một tập hợp của Route objects mà gắn Http Requests tới một hàm gọi.
- Thông thường thì Webservices có một root path (/user) và định nghĩa một kiểu MIME cho routes của nó.
- Webservices phải được added vào một container để xử lý Http request từ một server.
- Một Route được định nghĩa bởi một HTTP method, một URL path và tùy chọn kiểu MIME mà nó sử dụng (Content-Type) và produces (Accept).
- Trong repo này sử dụng Thư viện (package) go-restful có logic để tìm được Route tốt nhất và nếu tìm được sẽ gọi function tương ứng.

- Ví dụ khi tạo webservice và Route, gắn Route với action function tương ứng:
```
ws := new(restful.WebService)
ws.
	Path("/users").
	Consumes(restful.MIME_JSON, restful.MIME_XML).
	Produces(restful.MIME_JSON, restful.MIME_XML)

ws.Route(ws.GET("/{user-id}").To(u.findUser))  // u is a UserResource

...

// GET http://localhost:8080/users/1
func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	...
}
```
2. Regular expression matching routes
- Một Route parameter có thể được định nghĩa sử dụng format ```uri/{var[:regexp]}``` hoặc là một cấu trúc đặc biệt của nó ```uri/{var:*}``` cho việc matching phần đuôi của path.
- Ví dụ ```/persons/{name:{name:[A-Z][A-Z]}``` có thể được dùng để lấy giá trị của biến số name.
3. Containers
- Container giữ một tập của các Webservice, Filters và một http.ServeMux cho việc ghép kênh các http requests.
- Khi sử dụng package go-restful, sử dụng cấu trúc ```restful.Add(…)``` và ```restful.Filter()``` sẽ register Webservices và Filters vào một Default container.
- Default container của go-restful sử dụng http.DefaultServeMux
- Chúng ta có thể tự tạo một Container và tạo mới một http.Serve cho container đó:
```
container := restful.NewContainer()
ws := bookservice.NewAPIServer() 
container.Add(ws)
```

Run server with log:
```
log.Fatal(http.ListenAndServe("0.0.0.0:8888", container))
```


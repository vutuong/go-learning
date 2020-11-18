go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

or
If you are using Go modules:
Ensure your gRPC-Go version is ```required``` at the appropriate version in the same module containing the generated ```.pb.go``` files. For example, SupportPackageIsVersion6 needs ```v1.27.0```, so in your go.mod file:
```
module <your module name>

require (
    google.golang.org/grpc v1.27.0
)
```

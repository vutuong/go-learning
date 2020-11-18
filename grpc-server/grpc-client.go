package main
import (
	"log"
	"context"
	"google.golang.org/grpc"
	"example.com/grpc-server/tuong"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	t := tuong.NewTuongServiceClient(conn)

	message := tuong.Message{
		Body: "Hello from client !",
	}

	response, err := t.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from Server: %s", response.Body)
}
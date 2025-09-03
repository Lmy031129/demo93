package inits

import (
	"api-gatware/basic/globar"

	"flag"
	"log"

	__ "api-gatware/basic/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitGrpc() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	globar.Client = __.NewUserClient(conn)

}

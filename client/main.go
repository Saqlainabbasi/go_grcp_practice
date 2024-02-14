package main

import (
	"log"

	pb "github.com/Saqlainabbasi/go_grcp_practice/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// define a port constent
const (
	port = ":8080"
)

func main() {
	// make a connection
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection failed %v", err)
	}
	//close the connection
	defer conn.Close()
	// use the pb file for client side service initialization
	client := pb.NewGreetServiceClient(conn)

	//create a name list of user
	//use the pb file to make list
	names := &pb.NameList{
		Names: []string{"Bob", "Alice", "Jhon"},
	}
	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callHelloBidirectionalStream(client, names)
}

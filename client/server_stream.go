package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Saqlainabbasi/go_grcp_practice/proto"
)

// create a function
// which except client and pointer ti the name list
func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Println("Stream started")
	//call the server function from the client side
	//which takes the context and names of users
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	//check err
	if err != nil {
		log.Fatalf("Could not send strem: %v", err)
	}
	//for loop to read the message form the stream
	for {
		message, err := stream.Recv()
		//check if the end of the file is reached then break  the loop
		if err == io.EOF {
			break
		}
		log.Println(message)
	}
	log.Println("Sreaming finished")

}

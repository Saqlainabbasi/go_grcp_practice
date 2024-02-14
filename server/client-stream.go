package main

import (
	"io"
	"log"

	pb "github.com/Saqlainabbasi/go_grcp_practice/proto"
)

// create a func/method
// recieve pb client stream and return an error
// reciever func
func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	//make slice of string
	var message []string
	// make a infinate for loop for listening to the stream form the client
	for {
		req, err := stream.Recv()
		//if the end of the file is reached break the loop
		if err == io.EOF {
			//send the response back
			return stream.SendAndClose(&pb.MessageList{Messages: message})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request from: %v", req.Name)
		message = append(message, "Hello "+req.Name)
	}
	// print the user
	// append the slice
}

package main

import (
	"io"
	"log"

	pb "github.com/Saqlainabbasi/go_grcp_practice/proto"
)

//create func
//reviever func
//params pb and return error

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	// infinite for loop for recieving the req
	for {
		req, err := stream.Recv()
		//check pfor end of file
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("Got a request from %v", req.Name)
		//make a response
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		//send the response back....
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

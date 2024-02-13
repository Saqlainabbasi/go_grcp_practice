package main

import (
	"log"
	"time"

	pb "github.com/Saqlainabbasi/go_grcp_practice/proto"
)

// create method
// reciever func
// takes param like pointer to pb nameList and stream
// returns pointer pb Response and error
func (s *helloServer) callSayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	//print the name list
	log.Printf("got the req with names %v", req.Names)
	//loop over the name list
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
	//use time func to create a delay in res
	return nil
}

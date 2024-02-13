package main

import (
	"context"

	pb "github.com/Saqlainabbasi/go_grcp_practice/proto"
)

// make func
// reciever func of type helloServer type
// take in context,pointer no params from pb and pb response
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}

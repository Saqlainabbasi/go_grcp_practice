package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Saqlainabbasi/go_grcp_practice/proto"
)

// create a func
// recieves client and pointer to the name list
func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	//log
	log.Println("Client stream started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Chould not send names: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		//send the msg and check for the error
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sending req with name: %v", name)
		time.Sleep(3 * time.Second)
	}
	//check if the stream is end
	res, err := stream.CloseAndRecv()
	log.Println("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while recieving %v", err)
	}
	log.Println(res.Messages)
}

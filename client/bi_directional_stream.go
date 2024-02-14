package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Saqlainabbasi/go_grcp_practice/proto"
)

// make func
// recieve the client and nameList pointer
func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Println("Bidirectional Streaming started")
	//call the server func
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could send names %v", names)
	}
	//create a chanel
	// this chanel will waite for the go routine to finish
	done := make(chan struct{})
	//infinate the for loop to send the req
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(done)
	}()
	//range on the name list
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		log.Printf("Sending req with name %v", name)
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending req %v", err)
		}
		//sleep for delay in sending the req
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	//read the signal form the cahnel
	<-done
	println("Streaming finished")

}

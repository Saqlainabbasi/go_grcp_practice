package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Saqlainabbasi/go_grcp_practice/proto"
)

// make func
// take client as param
func callSayHello(client pb.GreetServiceClient) {

	//create context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// call func
	res, err := client.SayHello(ctx, &pb.NoParam{})
	// error check
	if err != nil {
		log.Fatalf("Could not greet %v", err)
	}
	//log the msg
	log.Printf("%s", res.Message)
}

package main

import (
	"context"
	"fmt"
	
	"log"

	pb "netxd_grpc_mongo/netxd_customer" 
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	

	response, err := client.CreateCustomer(context.Background(), &pb.Customer{Customer_ID : 50})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response.Customer_ID )
}
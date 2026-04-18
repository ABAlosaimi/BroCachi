package main

import (
	"context"
	"time"

	"github.com/ABAlosaimi/BroCachi/services/common/protoGen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := protoGen.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	res, err := client.CreateOrder(ctx, &protoGen.CreateOrderRequest{
		OrderId: "123",
		ProductId: "34567",
		Quantity: 3,
	})
	if err != nil {
		panic(err)
	}

	println("Order created with:", res.Message)
}
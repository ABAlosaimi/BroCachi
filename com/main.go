package main

import (
	"github.com/ABAlosaimi/BroCachi/services/orders"
)

func main() {
	grpcServer := orders.NewgRPCServer(":9000")
	if err := grpcServer.Start(); err != nil {
		panic(err)
	}
}
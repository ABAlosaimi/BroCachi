package grpc

import (
	"github.com/ABAlosaimi/BroCachi/services/common/protoGen"
	"github.com/ABAlosaimi/BroCachi/services/orders/types"
)

type OrderGRPCHandler struct {
	orderService types.OrderService
	protoGen.UnimplementedOrderServiceServer
}

func NewOrderGRPCHandler(orderService types.OrderService) *OrderGRPCHandler {
	return &OrderGRPCHandler{
		orderService: orderService,
	}
}

// implement the gRPC methods here

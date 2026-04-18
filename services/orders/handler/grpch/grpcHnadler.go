package grpch

import (
	"context"
	"github.com/ABAlosaimi/BroCachi/services/common/protoGen"
	"github.com/ABAlosaimi/BroCachi/services/orders/types"
	"google.golang.org/grpc"
)

// this equivalent to spring controller in java or handler in nodejs, it will handle the incoming gRPC requests and call the appropriate methods in the service layer
type OrderGRPCHandler struct {
	orderService types.OrderService
	protoGen.UnimplementedOrderServiceServer
}

func NewOrderGRPCHandler(grpc *grpc.Server,orderService types.OrderService) *OrderGRPCHandler {
	gRPChandler := &OrderGRPCHandler{
		orderService: orderService,
	}

	protoGen.RegisterOrderServiceServer(grpc, gRPChandler)

	return gRPChandler
}

// implement the gRPC methods here
// this is the handler that will be registered in the gRPC server and will handle the incoming requests and call the appropriate methods in the service layer
func (h *OrderGRPCHandler) CreateOrder(cxt context.Context, req *protoGen.CreateOrderRequest) (*protoGen.CreateOrderResponse, error) {
 	res,err := h.orderService.CreateOrder(cxt, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

package service

import (
	"context"

	"github.com/ABAlosaimi/BroCachi/services/common/protoGen"
	"github.com/ABAlosaimi/BroCachi/services/orders/types"
)

type OrdersService struct {
	orderService types.OrderService
}

func NewOrdersService() *OrdersService {
	return  &OrdersService{} 
}

func (s *OrdersService) CreateOrder(cxt context.Context, req *protoGen.CreateOrderRequest) (*protoGen.CreateOrderResponse, error) {
	return &protoGen.CreateOrderResponse{
		Message: "Order created successfully",
	}, nil
}
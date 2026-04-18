package service

import (
	"context"
	"github.com/ABAlosaimi/BroCachi/services/common/protoGen"
)

type OrdersService struct {

}

func NewOrdersService() *OrdersService {
	return  &OrdersService{}
}

func CreateOrder(cxt context.Context, req *protoGen.CreateOrderRequest) (*protoGen.CreateOrderResponse, error) {
	return &protoGen.CreateOrderResponse{
		Message: "Order created successfully",
	}, nil
}
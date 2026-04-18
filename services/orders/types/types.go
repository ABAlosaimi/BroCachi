package types

import (
	"context"
	"github.com/ABAlosaimi/BroCachi/services/common/protoGen"
)

type OrderService interface {
 CreateOrder(context.Context, *protoGen.CreateOrderRequest) (*protoGen.CreateOrderResponse, error)
}
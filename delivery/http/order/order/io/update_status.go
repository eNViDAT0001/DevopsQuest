package io

import "github.com/eNViDAT0001/GolangAdventure/internal/order/entities"

type UpdateOrderStatusReq struct {
	Status entities.OrderStatus `json:"status"`
}

package converter

import (
	"order/internal/model"
	desc "shared/pkg/proto/order/v1"
)

func ToDescFromOrderModel(order *model.Order) *desc.GetOrderResponse {
	if order == nil {
		return nil
	}
	return &desc.GetOrderResponse{
		OrderUuid:  order.UUID,
		Status:     order.Status,
		TotalPrice: float32(order.TotalPrice),
		Items:      ToDescFromOrderItems(order.Items),
	}
}

func ToDescFromOrderItems(items []model.OrderItem) []*desc.OrderItem {
	var res []*desc.OrderItem
	for _, item := range items {
		res = append(res, &desc.OrderItem{
			PartUuid: item.PartUUID,
			Quantity: uint32(item.Quantity),
		})
	}
	return res
}

func ToOrderFromDesc(req *desc.CreateOrderRequest) *model.Order {

	items := make([]model.OrderItem, 0, len(req.GetItems()))

	for _, item := range req.GetItems() {
		items = append(items, model.OrderItem{
			PartUUID: item.GetPartUuid(),
			Quantity: int32(item.GetQuantity()),
		})
	}

	return &model.Order{
		UserUUID: req.GetUserUuid(),
		Items:    items,
	}

}

func ToOrderCancelRequest(req *desc.CancelOrderRequest) string {
	return req.GetOrderUuid()
}

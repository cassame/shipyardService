package converter

import (
	"order/internal/model"
	openapi "shared/pkg/openapi/order/v1"
)

func ToOrderFromAPI(req *openapi.CreateOrderRequest) *model.Order {
	if req == nil {
		return nil
	}

	items := make([]model.OrderItem, 0, len(req.PartUuids))
	for _, partUUID := range req.PartUuids {
		items = append(items, model.OrderItem{
			PartUUID: partUUID,
			Quantity: 1,
		})
	}

	return &model.Order{
		UserUUID: req.UserUUID,
		Items:    items,
	}
}

func ToAPIFromOrderModel(order *model.Order) *openapi.Order {
	if order == nil {
		return nil
	}

	partUUIDs := make([]string, 0, len(order.Items))
	for _, item := range order.Items {
		partUUIDs = append(partUUIDs, item.PartUUID)
	}

	return &openapi.Order{
		UUID:      order.UUID,
		UserUUID:  order.UserUUID,
		PartUuids: partUUIDs,
		Status:    openapi.OrderStatus(order.Status),
	}
}

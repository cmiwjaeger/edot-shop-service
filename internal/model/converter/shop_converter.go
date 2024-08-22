package converter

import (
	"edot-monorepo/services/shop-service/internal/entity"
	"edot-monorepo/services/shop-service/internal/model"
	"edot-monorepo/shared/events"
)

func ShopToResponse(item *entity.Shop) *model.ShopCreateResponse {
	return &model.ShopCreateResponse{
		ID:        item.ID,
		Name:      item.Name,
		Address:   item.Address,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func ShopWhToResponse(item *entity.ShopWarehouse) *model.ShopAssignWarehouseResponse {
	return &model.ShopAssignWarehouseResponse{
		ShopID:      item.ShopID,
		WarehouseID: item.WarehouseID,
	}
}

func ShopToEvent(item *entity.Shop) *events.ShopCreatedEvent {
	return &events.ShopCreatedEvent{
		ID:        item.ID,
		Name:      item.Name,
		Address:   item.Address,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func ShopWhToEvent(item *entity.ShopWarehouse, assign bool) *events.ShopWarehouseAssignedEvent {
	return &events.ShopWarehouseAssignedEvent{
		Assigned:    assign,
		WarehouseID: item.WarehouseID,
		ShopID:      item.ShopID,
	}
}

func ShopToTokenResponse(user *entity.Shop) *model.ShopCreateResponse {
	return &model.ShopCreateResponse{}
}

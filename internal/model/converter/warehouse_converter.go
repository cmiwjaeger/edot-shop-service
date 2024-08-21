package converter

import (
	"edot-monorepo/services/shop-service/internal/entity"
	"edot-monorepo/services/shop-service/internal/model"
)

func ShopToResponse(user *entity.Shop) *model.ShopResponse {
	return &model.ShopResponse{
		ID: user.ID,

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ShopToTokenResponse(user *entity.Shop) *model.ShopResponse {
	return &model.ShopResponse{}
}

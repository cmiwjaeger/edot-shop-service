package route

import (
	http "edot-monorepo/services/shop-service/internal/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App *fiber.App

	ShopController *http.ShopController
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}
func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/shop", c.ShopController.Create)
	c.App.Post("/api/shop/assign-warehouse", c.ShopController.AssignWarehouse)

}

func (c *RouteConfig) SetupAuthRoute() {
	// c.App.Use(c.AuthMiddleware)
}

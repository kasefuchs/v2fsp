package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kasefuchs/v2fsp/internal/app/v2fsp/controller"
)

func SubscriptionRoute(app *fiber.App) {
	r := app.Group("/subscription")

	r.All("/", controller.SubscriptionGet)
}

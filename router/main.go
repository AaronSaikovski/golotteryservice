package router

import (
	"github.com/AaronSaikovski/golotteryservice/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/ping", handlers.HandleHealthCheck)
	app.Get("/lottery", handlers.HandleGetLotteryNumbers)
}

package routes

import (
	"Blug/api/handler"
	"Blug/middleware"
	"Blug/pkg/user"

	"github.com/gofiber/fiber/v2"
)

// UserRouter is the Router for GoFiber App
func UserRouter(app fiber.Router, service *user.Service) {
	// 公开路由（不需要认证）
	app.Post("/register", handler.Register(service))
	app.Post("/login", handler.Login(service))
	app.Get("/public/user", handler.GetUserByName(service))
	app.Get("/myInfo", handler.GetMyInfo(service))
	// 认证路由组
	auth := app.Group("/user", middleware.JwtAuthMiddleware())
	//authed ping check
	auth.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"mgs": "pong!"})
	})

	auth.Put("/update", handler.UpdateUser(service))
	auth.Delete("/delete", handler.DeleteUserById(service))
	auth.Get("/id", handler.GetUserById(service))
}

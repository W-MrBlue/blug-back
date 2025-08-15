package main

import (
	"Blug/api/routes"
	"Blug/config"
	"Blug/pkg/article"
	"Blug/pkg/db"
	"Blug/pkg/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"log"
)

func main() {
	newDb := db.SqliteInit()

	config.InitConfig()

	// Initialize article service
	articleRepository := article.NewRepository(newDb)
	articleService := article.NewService(*articleRepository)

	// Initialize user service
	userRepository := user.NewRepository(newDb)
	userService := user.NewService(*userRepository)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Send([]byte("pong"))
	})

	api := app.Group("/api")
	routes.ArticleRouter(api, articleService)
	routes.UserRouter(api, userService)
	log.Fatal(app.Listen(":8080"))
}

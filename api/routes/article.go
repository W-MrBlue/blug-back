package routes

import (
	"Blug/api/handler"
	"Blug/middleware"
	"Blug/pkg/article"

	"github.com/gofiber/fiber/v2"
)

// ArticleRouter is the Router for GoFiber App
func ArticleRouter(app fiber.Router, service *article.Service) {
	// 公开路由
	app.Get("/public/articles", handler.GetAllArticles(service))          // 获取所有文章
	app.Get("/public/articles/paged", handler.GetArticlesByPage(service)) // 分页获取
	app.Get("/public/articles/count", handler.GetArticleCount(service))
	app.Get("/public/article/:id", handler.GetArticle(service)) // 获取单篇文章
	app.Get("/public/classes", handler.GetAllClasses(service))

	// 认证路由组 (统一前缀)
	auth := app.Group("/article", middleware.JwtAuthMiddleware()) //middleware.JwtAuthMiddleware())
	auth.Post("/create", handler.AddArticle(service))             // POST /articles
	auth.Put("/update", handler.UpdateArticle(service))           // PUT /articles/123
	auth.Put("/delete", handler.DeleteArticle(service))           // DELETE /articles/123
	auth.Put("/delete_undo", handler.DeleteArticleUndo(service))  // DELETE /articles/123
}

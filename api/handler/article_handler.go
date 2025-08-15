package handler

import (
	"Blug/api/presenter"
	"Blug/pkg/article"
	"Blug/pkg/entities"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func AddArticle(service *article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusBadRequest))
		}
		result, err := service.AddArticle(requestBody.Title, requestBody.Content)
		if err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusInternalServerError))
		}
		return c.JSON(presenter.ArticleSuccessRespWithData(result, http.StatusOK))
	}
}

func GetArticle(service *article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idString := c.Params("id")
		if idString == "" {
			return c.JSON(presenter.UserErrorResp(fmt.Errorf("id empty"), http.StatusBadRequest))
		}
		id, err := strconv.Atoi(idString)
		if err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusBadRequest))
		}
		// 打印日志
		println("request article id: ", id)

		result, err := service.GetArticle(id)
		if err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusInternalServerError))
		}
		return c.JSON(presenter.ArticleSuccessRespWithData(result, http.StatusOK))
	}
}

func GetAllArticles(service *article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllArticles()
		if err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusInternalServerError))
		}
		return c.JSON(presenter.ArticleSuccessRespWithList(result, http.StatusOK))
	}
}

func GetArticleCount(service *article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result := service.GetArticleCount()
		return c.JSON(presenter.ArticleCountSuccess(result, http.StatusOK))
	}
}

func GetAllClasses(service *article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetAllClasses()
		if err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusInternalServerError))
		}
		return c.JSON(presenter.ClassSuccessRespWithList(result, http.StatusOK))
	}
}

func UpdateArticle(service *article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Id      int    `json:"id"`
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusBadRequest))
		}

		// 创建文章实体
		articleEntity := &entities.Article{
			Id:      requestBody.Id,
			Title:   requestBody.Title,
			Content: requestBody.Content,
		}

		// 更新文章
		result, err := service.UpdateArticle(articleEntity)
		if err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusInternalServerError))
		}

		// 返回更新后的数据
		return c.JSON(presenter.ArticleSuccessRespWithData(result, http.StatusOK))
	}
}

func DeleteArticle(service *article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Id int `json:"id"`
		}
		if err := c.QueryParser(&requestBody); err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusBadRequest))
		}

		// 打印日志
		println("request article id: ", requestBody.Id)
		err := service.DeleteArticle(requestBody.Id)
		if err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusInternalServerError))
		}
		return c.JSON(presenter.ArticleSuccessResp(http.StatusOK))
	}
}
func DeleteArticleUndo(service *article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Id int `json:"id"`
		}
		if err := c.QueryParser(&requestBody); err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusBadRequest))
		}

		// 打印日志
		println("request article id: ", requestBody.Id)
		err := service.DeleteArticleUndo(requestBody.Id)
		if err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusInternalServerError))
		}
		return c.JSON(presenter.ArticleSuccessResp(http.StatusOK))
	}
}

func GetArticlesByPage(service *article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Page        int  `query:"page"`
			Limit       int  `query:"limit"`
			ShowDeleted bool `query:"show_deleted"`
		}
		if err := c.QueryParser(&requestBody); err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusBadRequest))
		}

		println("request page: ", requestBody.Page)
		println("show deleted: ", requestBody.ShowDeleted)
		result, err := service.GetArticlesByPage(requestBody.Page, requestBody.Limit, requestBody.ShowDeleted)
		if err != nil {
			return c.JSON(presenter.ArticleErrorResp(err, http.StatusInternalServerError))
		}
		return c.JSON(presenter.ArticleSuccessRespWithList(result, http.StatusOK))
	}
}

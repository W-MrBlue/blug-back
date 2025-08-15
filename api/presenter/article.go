package presenter

import (
	"Blug/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type Article struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func ArticleSuccessRespWithData(data *entities.Article, code int) *fiber.Map {
	return &fiber.Map{
		"code":  code,
		"data":  data,
		"error": nil,
	}
}
func ArticleCountSuccess(cnt, code int64) *fiber.Map {
	return &fiber.Map{
		"code":  code,
		"data":  cnt,
		"error": nil,
	}
}

func ArticleSuccessRespWithList(data []*entities.Article, code int) *fiber.Map {
	return &fiber.Map{
		"code":  code,
		"data":  data,
		"error": nil,
	}
}

func ArticleSuccessResp(code int) *fiber.Map {
	return &fiber.Map{
		"code":  code,
		"data":  "this is a success response,but without data,so I write this to fill it",
		"error": nil,
	}
}

func ArticleErrorResp(err error, code int) *fiber.Map {
	return &fiber.Map{
		"code":  code,
		"data":  "this is a error response,but without data,so I write this to fill it",
		"error": err.Error(),
	}
}

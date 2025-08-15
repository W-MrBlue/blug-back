package presenter

import (
	"Blug/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type class struct {
	Name string `json:"name"`
}

func newClass(data *entities.Class) *class {
	return &class{
		Name: data.ClassName,
	}
}

func newClassGroup(data []*entities.Class) []*class {
	var result []*class
	for _, cls := range data {
		result = append(result, newClass(cls))
	}
	return result
}

func ClassSuccessRespWithList(data []*entities.Class, code int) *fiber.Map {
	return &fiber.Map{
		"code": code,
		"data": newClassGroup(data),
	}
}

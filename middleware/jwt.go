package middleware

import (
	"Blug/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"net/http"
)

func JwtAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := utils.ExtractTokenID(c)
		if err != nil {
			return c.JSON(&fiber.Map{
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
		}
		c.Set("userId", strconv.Itoa(id))
		return c.Next()
	}
}

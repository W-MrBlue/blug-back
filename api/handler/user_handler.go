package handler

import (
	"Blug/api/presenter"
	"Blug/pkg/user"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AddUser(service *user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusBadRequest))
		}
		result, err := service.AddUser(requestBody.Name, requestBody.Password)
		if err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusInternalServerError))
		}
		return c.JSON(presenter.UserSuccessRespWithData(result, http.StatusOK))
	}
}

func GetUserById(service *user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Id int `json:"id"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusBadRequest))
		}

		result, err := service.GetUserById(requestBody.Id)
		if err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusInternalServerError))
		}
		return c.JSON(presenter.UserSuccessRespWithData(result, http.StatusOK))
	}
}

func GetUserByName(service *user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Name string `json:"name"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusBadRequest))
		}
		result, err := service.GetUserByName(requestBody.Name)
		if err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusInternalServerError))
		}
		return c.JSON(presenter.UserSuccessRespWithData(result, http.StatusOK))
	}
}

func UpdateUser(service *user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusBadRequest))
		}

		result, err := service.UpdateUser(requestBody.Name, requestBody.Password)
		if err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusInternalServerError))
		}

		return c.JSON(presenter.UserSuccessRespWithData(result, http.StatusOK))
	}
}

func DeleteUserById(service *user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Id int `json:"id"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusBadRequest))
		}

		err := service.DeleteUserById(requestBody.Id)
		if err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusInternalServerError))
		}
		return c.JSON(presenter.UserSuccessResp(http.StatusOK))
	}
}

func Login(service *user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return err
		}
		token, err := service.CheckPassword(requestBody.Username, requestBody.Password)
		if err != nil {
			return c.JSON(presenter.LoginErrorResp(err, http.StatusUnauthorized))
		} else {
			return c.JSON(presenter.LoginSuccessResp(token, http.StatusOK))
		}
	}
}

func Register(service *user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return err
		}
		res, err := service.AddUser(requestBody.Username, requestBody.Password)
		if err != nil {
			return c.JSON(presenter.RegisterErrorResp(err, http.StatusInternalServerError))
		} else {
			token, err := service.CheckPassword(res.Name, res.Password)
			if err != nil {
				return c.JSON(presenter.RegisterErrorResp(err, http.StatusUnauthorized))
			} else {
				return c.JSON(presenter.RegisterSuccessResp(token, http.StatusOK))
			}
		}
	}
}

func GetMyInfo(service *user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := 1
		res, err := service.GetUserById(id)
		if err != nil {
			return c.JSON(presenter.UserErrorResp(err, http.StatusInternalServerError))
		} else {
			return c.JSON(presenter.UserSuccessRespWithData(res, http.StatusOK))
		}
	}
}

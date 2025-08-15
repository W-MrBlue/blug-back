package presenter

import (
	"Blug/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Signature string `json:"signature"`
	PFPUrl    string `json:"pfpUrl"`
}

type LoginResponse struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}

type RegisterResponse struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}

func UserSuccessRespWithData(data *entities.User, code int) *fiber.Map {
	return &fiber.Map{
		"code": code,
		"data": &User{
			Id:        data.Id,
			Name:      data.Name,
			PFPUrl:    data.PFPUrl,
			Signature: data.Signature,
		},
		"error": nil,
	}
}

func UserSuccessResp(code int) *fiber.Map {
	return &fiber.Map{
		"code":  code,
		"error": nil,
	}
}

func UserErrorResp(err error, code int) *fiber.Map {
	return &fiber.Map{
		"code":  code,
		"data":  "",
		"error": err.Error(),
	}
}

func LoginSuccessResp(token string, code int) *fiber.Map {
	return &fiber.Map{
		"code":  code,
		"token": token,
	}
}

func LoginErrorResp(err error, code int) *fiber.Map {
	return &fiber.Map{
		"code":  code,
		"error": err.Error(),
	}
}

func RegisterSuccessResp(token string, code int) *fiber.Map {
	return &fiber.Map{
		"code":  code,
		"token": token,
	}
}

func RegisterErrorResp(err error, code int) *fiber.Map {
	return &fiber.Map{
		"code":  code,
		"error": err.Error(),
	}
}

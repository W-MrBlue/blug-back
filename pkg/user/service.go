package user

import (
	"Blug/pkg/entities"
	"Blug/utils"
	"fmt"
)

type ServiceInterface interface {
	AddUser(userName, passWord string) (*entities.User, error)
	GetUserById(id int) (*entities.User, error)
	GetUserByName(name string) (*entities.User, error)
	UpdateUser(userName, passWord string) (*entities.User, error)
	DeleteUserById(id int) error
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) AddUser(userName, passWord string) (*entities.User, error) {
	return s.repository.AddUser(userName, passWord)
}

func (s *Service) GetUserById(id int) (*entities.User, error) {
	return s.repository.GetUserById(id)
}

func (s *Service) GetUserByName(name string) (*entities.User, error) {
	return s.repository.GetUserByName(name)
}

func (s *Service) UpdateUser(userName, passWord string) (*entities.User, error) {
	return s.repository.UpdateUser(userName, passWord)
}

func (s *Service) DeleteUserById(id int) error {
	return s.repository.DeleteUserById(id)
}

func (s *Service) CheckPassword(userName, passWord string) (string, error) {
	user, err := s.repository.GetUserByName(userName)
	if err != nil || user.Password != passWord {
		return "", fmt.Errorf("用户不存在或密码错误")
	}
	// 生成JWT token
	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

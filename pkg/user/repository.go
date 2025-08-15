package user

import (
	"Blug/pkg/entities"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Db: db,
	}
}

func (r *Repository) AddUser(userName, passWord string) (*entities.User, error) {
	user := &entities.User{
		Name:     userName,
		Password: passWord,
	}
	if err := r.Db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserById(id int) (*entities.User, error) {
	user := &entities.User{}
	if err := r.Db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserByName(name string) (*entities.User, error) {
	user := &entities.User{}
	if err := r.Db.Where("name = ?", name).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) UpdateUser(userName, passWord string) (*entities.User, error) {
	user := &entities.User{}
	if err := r.Db.Where("name = ?", userName).First(user).Error; err != nil {
		return nil, err
	}

	if err := r.Db.Model(user).Updates(map[string]interface{}{
		"Password": passWord,
	}).Error; err != nil {
		return nil, err
	}

	// 获取更新后的用户信息
	updatedUser := &entities.User{}
	if err := r.Db.First(updatedUser, user.Id).Error; err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *Repository) DeleteUserById(id int) error {
	if err := r.Db.Delete(&entities.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

package user

import (
	"application/internal/model"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) Add(dto *CreateUserDTO) (model.User, error) {
	user := dto.toModel()

	err := s.db.Create(&user).Error
	return user, err
}

func (s *UserService) Get() ([]model.User, error) {
	users := []model.User{}

	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

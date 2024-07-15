package user

import (
	"application/internal/model"
	"time"

	"github.com/jinzhu/copier"
)

type CreateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (dto *CreateUserDTO) toModel() model.User {
	user := model.User{}
	copier.Copy(&user, dto)
	return user
}

type UserDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (UserDTO) fromModelArr(users []model.User) []UserDTO {
	dto := []UserDTO{}
	copier.Copy(&dto, &users)
	return dto
}
func (dto UserDTO) fromModel(users model.User) UserDTO {
	copier.Copy(&dto, &users)
	return dto
}

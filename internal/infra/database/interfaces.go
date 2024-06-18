package database

import "github.com/thiagohmm/ApiGo/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaild string) (*entity.User, error)
}

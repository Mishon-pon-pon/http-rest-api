package store

import "github.com/mishon-pon-pon/http-rest-api/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}

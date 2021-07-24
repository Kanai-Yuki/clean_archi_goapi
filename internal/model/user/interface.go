package user

import "github.com/Kanai-Yuki/clean_archi_goapi/internal/entities"

type InterfaceModel interface {
	// User
	GetUser(string) (*entities.User, error)
	CreateUser(*entities.User) (string, error)
}

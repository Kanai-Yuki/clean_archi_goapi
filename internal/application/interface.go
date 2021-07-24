package application

import "github.com/Kanai-Yuki/clean_archi_goapi/internal/entities"

// InterfaceApplication ...
// Input Boundary & Output Boundary
type InterfaceApplication interface {
	// User
	GetUser(uuid string) (*entities.User, error)
	CreateUser(name string, age int64) (string, error)
}

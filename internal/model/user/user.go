package user

import (
	"fmt"

	"github.com/Kanai-Yuki/clean_archi_goapi/internal/entities"
)

// CreateUser...
func (m Model) CreateUser(u *entities.User) (string, error) {
	fmt.Println("CreateUserまできてるよ")
	return "", nil
}

// GetUser...
func (m Model) GetUser(uuid string) (*entities.User, error) {
	return &entities.User{}, nil
}

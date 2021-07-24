package application

import (
	"github.com/Kanai-Yuki/clean_archi_goapi/internal/entities"

	"github.com/google/uuid"
)

// Use Case Interactor
// ドメイン(Entities)を生成し、DataAccessInterfaceを呼び出す
func (a Application) CreateUser(name string, age int64) (string, error) {
	// UseCaseInteractor(Application) → DataAccessInterface
	return a.InterfaceModel.CreateUser(createUserEntitiy(name, age))
}

func (a Application) GetUser(uuid string) (*entities.User, error) {
	return a.InterfaceModel.GetUser(uuid)
}

func createUserEntitiy(name string, age int64) *entities.User {
	uuid := uuid.New()

	return &entities.User{
		UUID: uuid.String(),
		Name: name,
		Age:  age,
	}
}

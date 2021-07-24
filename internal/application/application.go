package application

import "github.com/Kanai-Yuki/clean_archi_goapi/internal/model/user"

type Application struct {
	user.InterfaceModel
}

func New() Application {
	return Application{}
}

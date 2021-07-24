package application

import (
	modelUser "github.com/Kanai-Yuki/clean_archi_goapi/internal/model/user"
)

type Application struct {
	model modelUser.InterfaceModel
}

func New(model modelUser.InterfaceModel) *Application {
	return &Application{model: model}
}

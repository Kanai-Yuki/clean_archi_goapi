package user

import "github.com/Kanai-Yuki/clean_archi_goapi/internal/model"

type Model struct {
	bm model.BaseModel
}

func New() *Model {
	return &Model{}
}

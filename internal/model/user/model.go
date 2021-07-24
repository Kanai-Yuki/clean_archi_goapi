package user

import (
	"github.com/Kanai-Yuki/clean_archi_goapi/internal/model"
	"github.com/Kanai-Yuki/clean_archi_goapi/internal/postgres"
)

type Model struct {
	model.BaseModel
}

func New(posCli *postgres.Client) *Model {
	m := &Model{}
	m.SetPosCli(posCli.DB)
	return m
}

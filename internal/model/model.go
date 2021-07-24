package model

import (
	"database/sql"
)

type BaseModel struct {
	posCli *sql.DB
}

func (bm *BaseModel) GetPosCli() *sql.DB {
	return bm.posCli
}

func (bm *BaseModel) SetPosCli(posCli *sql.DB) {
	bm.posCli = posCli
}

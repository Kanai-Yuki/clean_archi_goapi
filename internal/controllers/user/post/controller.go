package post

import (
	"encoding/json"
	"net/http"

	"github.com/Kanai-Yuki/clean_archi_goapi/internal/application"
)

type Controller struct {
	app application.InterfaceApplication
}

func New(app application.InterfaceApplication) Controller {
	return Controller{app: app}
}

func (c Controller) Exec(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	// 入力値をアプリケーションが必要な値(InputData)に変換
	var req RequestPostUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, http.StatusBadRequest, err
	}

	// バリデーション

	// Controller to UseCaseInteractor
	// 変換した値をUseCaseInteractor(Applicatin)に渡す
	uuid, err := c.app.CreateUser(req.Name, req.Age)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// WEBではPresenterを呼び出さずOutputDataを返却
	return ResponsePostUser{UUID: uuid}, http.StatusOK, nil
}

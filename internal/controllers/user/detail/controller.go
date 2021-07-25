package detail

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
	var req RequestDetailUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, http.StatusBadRequest, err
	}

	// バリデーションチェック

	user, err := c.app.GetUser(req.UUID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return ResponseDetailUser{user.Name, user.Age}, http.StatusOK, nil
}

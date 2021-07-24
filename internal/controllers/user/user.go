package user

import (
	"net/http"

	"github.com/Kanai-Yuki/clean_archi_goapi/internal/application"
	"github.com/Kanai-Yuki/clean_archi_goapi/internal/controllers"
	"github.com/Kanai-Yuki/clean_archi_goapi/internal/controllers/user/post"
)

type Controller struct {
	application.InterfaceApplication
}

func New() Controller {
	return Controller{}
}

// PostUsers ...
func (c Controller) PostUser(w http.ResponseWriter, r *http.Request) {
	controller := post.New(c.InterfaceApplication)
	controllers.ResponseHandler(w, r, controller.Exec)
}

// DetailUsers ...
func (c Controller) DetailUser(w http.ResponseWriter, r *http.Request) {}

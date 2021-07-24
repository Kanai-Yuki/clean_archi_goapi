package user

import (
	"net/http"

	"github.com/Kanai-Yuki/clean_archi_goapi/internal/application"
	"github.com/Kanai-Yuki/clean_archi_goapi/internal/controllers"
	"github.com/Kanai-Yuki/clean_archi_goapi/internal/controllers/user/post"
)

type Controller struct {
	app application.InterfaceApplication
}

func New(app application.InterfaceApplication) *Controller {
	return &Controller{app: app}
}

// PostUsers ...
func (c Controller) PostUser(w http.ResponseWriter, r *http.Request) {
	controller := post.New(c.app)
	controllers.ResponseHandler(w, r, controller.Exec)
}

// DetailUsers ...
func (c Controller) DetailUser(w http.ResponseWriter, r *http.Request) {}

package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/Kanai-Yuki/clean_archi_goapi/internal/application"
	"github.com/Kanai-Yuki/clean_archi_goapi/internal/controllers/user"
	modelUser "github.com/Kanai-Yuki/clean_archi_goapi/internal/model/user"
	"github.com/Kanai-Yuki/clean_archi_goapi/internal/postgres"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	applyRoutes(r)

	return r
}

func applyRoutes(r *chi.Mux) {
	// healthcheck
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })

	// Connect Postgres
	posCli, err := postgres.New()
	if err != nil {
		log.Println(err)
	}

	r.Route("/users", func(r chi.Router) {
		ui := user.New(application.New(modelUser.New(posCli)))
		(func(ui user.InterfaceUser, r chi.Router) {
			// ユーザー新規登録 (POST /users)
			r.Post("/", ui.PostUser)
			// ユーザー情報取得 (GET /users/{uuid})
			r.Get("/{uuid}", ui.DetailUser)
		})(ui, r)
	})
}

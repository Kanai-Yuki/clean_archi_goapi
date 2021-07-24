package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/Kanai-Yuki/clean_archi_goapi/internal/controllers/user"
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

	r.Route("/users", func(r chi.Router) {
		ui := user.New()
		(func(ui user.InterfaceUser, r chi.Router) {
			// ユーザー新規登録 (POST /users)
			r.Post("/", ui.PostUser)
			// ユーザー情報取得 (GET /users/{uuid})
			r.Get("/{uuid}", ui.DetailUser)
		})(ui, r)
	})
}

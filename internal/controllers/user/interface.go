package user

import "net/http"

type InterfaceUser interface {
	// ユーザー新規登録API (POST /users)
	PostUser(w http.ResponseWriter, r *http.Request)
	// ユーザー情報取得API (GET /users/{uuid})
	DetailUser(w http.ResponseWriter, r *http.Request)
}

package middlewares

import (
	"focusflow/internals/auth"
	"focusflow/internals/utils"
	"net/http"
)

// Проверка аутентификации
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := auth.Store.Get(r, "session")
		authenticated, ok := session.Values["authenticated"].(bool)

		if !ok || !authenticated {
			http.Error(w, "Доступ запрещён", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Проверка прав администратора
func RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := auth.Store.Get(r, "session")
		authenticated, ok := session.Values["authenticated"].(bool)
		role, _ := session.Values["role"].(string)

		if !ok || !authenticated || (role != "admin" && role != "superadmin") {
			http.Error(w, "Доступ разрешён только администраторам", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Проверка прав суперадмина
func RequireSuperadmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := utils.GetCurrentUser(r)
		if err != nil || user.Role != "superadmin" {
			http.Error(w, "Только для суперадмина", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

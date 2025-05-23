package handlers

import (
	"encoding/json"
	"focusflow/internals/auth"
	"focusflow/internals/services"
	"focusflow/internals/utils"
	"net/http"
)

// Проверка на роль админа
func isAdmin(r *http.Request) bool {
	session, _ := auth.Store.Get(r, "session")
	role, _ := session.Values["role"].(string)
	return role == "admin" || role == "superadmin"
}

// Получение всех пользователей
func AdminGetUsers(w http.ResponseWriter, r *http.Request) {
	if !isAdmin(r) {
		utils.RespondWithError(w, http.StatusForbidden, "Только для администраторов")
		return
	}

	users, err := services.GetAllUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка получения списка пользователей")
		return
	}

	// Преобразовать пользователей в безопасный формат без паролей
	type SafeUser struct {
		ID             uint   `json:"id"`
		Username       string `json:"username"`
		Role           string `json:"role"`
		IsBlocked      bool   `json:"is_blocked"`
		ProfilePicture string `json:"profile_picture"`
	}

	var safeUsers []SafeUser
	for _, u := range users {
		safeUsers = append(safeUsers, SafeUser{
			ID:             u.ID,
			Username:       u.Username,
			Role:           u.Role,
			IsBlocked:      u.IsBlocked,
			ProfilePicture: u.ProfilePicture,
		})
	}

	utils.RespondWithJSON(w, http.StatusOK, safeUsers)
}

// Админ добавляет нового пользователя
func AdminAddUser(w http.ResponseWriter, r *http.Request) {
	if !isAdmin(r) {
		utils.RespondWithError(w, http.StatusForbidden, "Только для администраторов")
		return
	}

	var input struct {
		Username string `json:"username"`
		Password string `json:"password_hash"`
		Role     string `json:"role"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	err := services.CreateUser(services.CreateUserInput{
		Username: input.Username,
		Password: input.Password,
		Role:     input.Role,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Пользователь успешно добавлен"})
}

// Админ блокирует пользователя
func AdminBlockUser(w http.ResponseWriter, r *http.Request) {
	if !isAdmin(r) {
		utils.RespondWithError(w, http.StatusForbidden, "Только для администраторов")
		return
	}

	username := r.URL.Query().Get("username")
	err := services.BlockUser(username)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Пользователь заблокирован"})
}

// Админ разблокирует пользователя
func AdminUnblockUser(w http.ResponseWriter, r *http.Request) {
	if !isAdmin(r) {
		utils.RespondWithError(w, http.StatusForbidden, "Только для администраторов")
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Имя пользователя не указано")
		return
	}

	err := services.UnblockUser(username)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Пользователь разблокирован"})
}

// Админ удаляет пользователя
func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	if !isAdmin(r) {
		utils.RespondWithError(w, http.StatusForbidden, "Только для администраторов")
		return
	}

	username := r.URL.Query().Get("username")
	err := services.DeleteUser(username)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Пользователь удалён"})
}

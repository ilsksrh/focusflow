package handlers

import (
	"encoding/json"
	"focusflow/internals/services"
	"focusflow/internals/utils"
	"net/http"
)

// Регистрация пользователя
func Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	err := services.RegisterUser(services.RegisterInput{
		Username: input.Username,
		Password: input.Password,
		Role:     input.Role,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Пользователь зарегистрирован"})
}

// Вход пользователя
func Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	err := services.LoginUser(w, r, services.LoginInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Успешный вход"})
}

// Выход пользователя
func Logout(w http.ResponseWriter, r *http.Request) {
	err := services.LogoutUser(w, r)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка выхода")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Вы вышли из системы"})
}

// Получение информации о текущем пользователе
func Me(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Ошибка получения пользователя")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"username":        user.Username,
		"role":            user.Role,
		"profile_picture": user.ProfilePicture,
		"is_blocked":      user.IsBlocked,
	})
}

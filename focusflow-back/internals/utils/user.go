package utils

import (
	"focusflow/internals/auth"
	"focusflow/internals/db"
	"focusflow/internals/models"
	"net/http"
	"errors"
)

// Получение текущего пользователя по сессии
func GetCurrentUser(r *http.Request) (models.User, error) {
	session, _ := auth.Store.Get(r, "session")
	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		return models.User{}, errors.New("Пользователь не найден в сессии")
	}

	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, errors.New("Пользователь не найден в базе данных")
	}

	return user, nil
}

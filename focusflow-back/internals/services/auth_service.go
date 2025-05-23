package services

import (
	"errors"
	"focusflow/internals/auth"
	"focusflow/internals/db"
	"focusflow/internals/models"
	"focusflow/internals/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Структура для входа/регистрации
type RegisterInput struct {
	Username string
	Password string
	Role     string
}

type LoginInput struct {
	Username string
	Password string
}

// Регистрация нового пользователя
func RegisterUser(input RegisterInput) error {
	if err := utils.ValidateUsername(input.Username); err != nil {
		return err
	}
	if err := utils.ValidatePassword(input.Password); err != nil {
		return err
	}
	if err := utils.ValidateRole(input.Role); err != nil {
		return err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Ошибка хэширования пароля")
	}

	role := input.Role
	if role == "" {
		role = "user"
	}

	user := models.User{
		Username:     input.Username,
		PasswordHash: string(hashed),
		Role:         role,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return errors.New("Пользователь уже существует")
	}

	return nil
}

// Вход пользователя
func LoginUser(w http.ResponseWriter, r *http.Request, input LoginInput) error {
	var user models.User
	if err := db.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return errors.New("Неверное имя пользователя")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return errors.New("Неверный пароль")
	}

	session, _ := auth.Store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["username"] = user.Username
	session.Values["role"] = user.Role
	session.Save(r, w)

	return nil
}

// Выход пользователя
func LogoutUser(w http.ResponseWriter, r *http.Request) error {
	session, _ := auth.Store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	return session.Save(r, w)
}

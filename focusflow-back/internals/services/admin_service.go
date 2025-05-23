package services

import (
	"errors"
	"focusflow/internals/db"
	"focusflow/internals/models"
	"focusflow/internals/utils"
	"golang.org/x/crypto/bcrypt"
)

// Структура для создания пользователя
type CreateUserInput struct {
	Username string
	Password string
	Role     string
}

// Получение всех пользователей
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Добавление нового пользователя
func CreateUser(input CreateUserInput) error {
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

	user := models.User{
		Username:     input.Username,
		PasswordHash: string(hashed),
		Role:         input.Role,
	}

	return db.DB.Create(&user).Error
}

// Блокировка пользователя
func BlockUser(username string) error {
	if username == "" {
		return errors.New("Имя пользователя не указано")
	}
	return db.DB.Model(&models.User{}).Where("username = ?", username).Update("is_blocked", true).Error
}

func UnblockUser(username string) error {
	if username == "" {
		return errors.New("Имя пользователя не указано")
	}
	return db.DB.Model(&models.User{}).Where("username = ?", username).Update("is_blocked", false).Error
}
// Удаление пользователя
func DeleteUser(username string) error {
	if username == "" {
		return errors.New("Имя пользователя не указано")
	}
	return db.DB.Where("username = ?", username).Delete(&models.User{}).Error
}

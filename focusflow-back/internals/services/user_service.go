package services

import (
	"errors"
	"focusflow/internals/db"
	"focusflow/internals/models"
	"focusflow/internals/utils"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"path/filepath"
)

// Структуры для запросов
type UpdateProfileInput struct {
	NewUsername    string
	ProfilePicture string
}

type ChangePasswordInput struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}


// Обновление профиля пользователя
func UpdateUserProfile(user *models.User, input UpdateProfileInput) error {
	if input.NewUsername != "" {
		if err := utils.ValidateUsername(input.NewUsername); err != nil {
			return err
		}
		user.Username = input.NewUsername
	}

	if input.ProfilePicture != "" {
		user.ProfilePicture = input.ProfilePicture
	}

	return db.DB.Save(user).Error
}


func ChangeUserPassword(user *models.User, input ChangePasswordInput) error {
	// Проверка старого пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.OldPassword)); err != nil {
		return errors.New("Неверный текущий пароль")
	}

	// Проверка валидности нового пароля
	if err := utils.ValidatePassword(input.NewPassword); err != nil {
		return err
	}

	// Проверка на совпадение со старым паролем
	if input.OldPassword == input.NewPassword {
		return errors.New("Новый пароль не должен совпадать с текущим")
	}

	// Хэшируем новый пароль
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Ошибка хэширования пароля")
	}

	user.PasswordHash = string(hashed)
	return db.DB.Save(user).Error
}



// Загрузка фото профиля
func UploadUserProfilePicture(user *models.User, file io.Reader, filename string) (string, error) {
	os.MkdirAll("uploads", os.ModePerm)

	filePath := filepath.Join("uploads", filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	user.ProfilePicture = filePath
	if err := db.DB.Save(user).Error; err != nil {
		return "", err
	}

	return filePath, nil
}

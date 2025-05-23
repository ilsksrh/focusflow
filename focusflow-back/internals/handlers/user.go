package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"focusflow/internals/db"
	"focusflow/internals/services"
	"focusflow/internals/utils"
)

// Обновление профиля пользователя
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Ошибка получения пользователя")
		return
	}

	var input services.UpdateProfileInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный JSON")
		return
	}

	if err := services.UpdateUserProfile(&user, input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Профиль обновлён"})
}

// Смена пароля
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Ошибка получения пользователя")
		return
	}

	var input services.ChangePasswordInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}

	if err := services.ChangeUserPassword(&user, input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error()) // 👈 отправляем конкретную ошибку
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Пароль успешно изменён"})
}


// Загрузка фотографии профиля
func UploadProfilePicture(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Не авторизован")
		return
	}

	// Ограничение размера файла
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Ошибка при обработке формы")
		return
	}

	file, handler, err := r.FormFile("profile_picture")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Файл не найден")
		return
	}
	defer file.Close()

	// Убедимся, что папка uploads существует
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		if err := os.Mkdir("uploads", os.ModePerm); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Не удалось создать папку uploads")
			return
		}
	}

	// Сохраняем файл
	filename := fmt.Sprintf("uploads/%s", filepath.Base(handler.Filename))
	dst, err := os.Create(filename)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка сохранения файла")
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка копирования файла")
		return
	}

	// Обновляем путь к картинке у пользователя
	user.ProfilePicture = filename
	if err := db.DB.Save(&user).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка сохранения пользователя")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Фото профиля обновлено",
		"path":    filename,
	})
}

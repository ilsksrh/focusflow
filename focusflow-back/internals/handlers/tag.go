package handlers

import (
	"encoding/json"
	"focusflow/internals/models"
	"focusflow/internals/services"
	"focusflow/internals/utils"
	"net/http"
)

func CreateTag(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil || user.Role != "superadmin" {
		utils.RespondWithError(w, http.StatusForbidden, "Нет доступа")
		return
	}

	var tag models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}

	createdTag, err := services.CreateTag(&tag)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Не удалось создать тег")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdTag)
}

func UpdateTag(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDParam(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Некорректный ID")
		return
	}

	var updatedTag models.Tag
	if err := json.NewDecoder(r.Body).Decode(&updatedTag); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}

	updated, err := services.UpdateTag(uint(id), updatedTag.Name)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, updated)
}


func GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := services.GetTags()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Не удалось получить теги")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tags)
}

func GetTagsByTaskID(w http.ResponseWriter, r *http.Request) {
	taskID, err := utils.GetIDParam(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Некорректный ID задачи")
		return
	}

	tags, err := services.GetTagsByTaskID(uint(taskID))
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tags)
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil || user.Role != "superadmin" {
		utils.RespondWithError(w, http.StatusForbidden, "Нет доступа")
		return
	}
	id, err := utils.GetIDParam(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Некорректный ID")
		return
	}

	if err := services.DeleteTag(uint(id)); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Не удалось удалить тег")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Тег удалён"})
}

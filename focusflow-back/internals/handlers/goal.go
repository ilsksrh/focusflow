package handlers

import (
	"encoding/json"
    "focusflow/internals/auth"
    "focusflow/internals/db"  
	"focusflow/internals/models"
	"focusflow/internals/services"
	"focusflow/internals/utils"
	"net/http"
)

func CreateGoal(w http.ResponseWriter, r *http.Request) {
    var goal models.Goal
    if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
        return
    }

    session, _ := auth.Store.Get(r, "session")
    username, ok := session.Values["username"].(string)
    if !ok || username == "" {
        utils.RespondWithError(w, http.StatusUnauthorized, "Пользователь не найден")
        return
    }


    var user models.User
    if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка поиска пользователя")
        return
    }

    goal.UserID = user.ID

    if goal.Date == "" {
        utils.RespondWithError(w, http.StatusBadRequest, "Дата обязательна")
        return
    }

    if err := services.CreateGoal(&goal); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Не удалось создать цель")
        return
    }

    utils.RespondWithJSON(w, http.StatusCreated, goal)
}




func GetGoals(w http.ResponseWriter, r *http.Request) {
    user, err := utils.GetCurrentUser(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Не авторизован")
        return
    }

    goals, err := services.GetGoalsByUser(user.ID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Не удалось получить цели")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, goals)
}

func UpdateGoal(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDParam(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Некорректный ID")
		return
	}

	var updatedGoal models.Goal
	if err := json.NewDecoder(r.Body).Decode(&updatedGoal); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}

	goal, err := services.UpdateGoal(uint(id), updatedGoal)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, goal)
}

func DeleteGoal(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDParam(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Некорректный ID")
		return
	}

	user, err := utils.GetCurrentUser(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Не авторизован")
		return
	}

	var goal models.Goal
	if err := db.DB.First(&goal, id).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Цель не найдена")
		return
	}

	if goal.UserID != user.ID {
		utils.RespondWithError(w, http.StatusForbidden, "Вы не можете удалить чужую цель")
		return
	}

	if err := db.DB.Delete(&goal).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка при удалении")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Цель успешно удалена"})
}



package handlers

import (
	"encoding/json"
	"focusflow/internals/db" 
	"focusflow/internals/models"
	"focusflow/internals/services"
	"focusflow/internals/utils"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Неверный JSON")
        return
    }

    user, err := utils.GetCurrentUser(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Не авторизован")
        return
    }

    task.UserID = user.ID 
	task.TeamID = user.TeamID

    if err := services.CreateTask(&task); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка при создании задачи")
        return
    }
	db.DB.Preload("Tags").First(&task, task.ID)

    utils.RespondWithJSON(w, http.StatusCreated, task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
    user, err := utils.GetCurrentUser(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Не авторизован")
        return
    }

    tasks, err := services.GetTasksByUser(user.ID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Не удалось получить задачи")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, tasks)
}


func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDParam(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Некорректный ID")
		return
	}

	task, err := services.GetTaskByID(uint(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, task)
}

func GetTasksByGoalID(w http.ResponseWriter, r *http.Request) {
	goalID, err := utils.GetIDParam(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Некорректный goal_id")
		return
	}

	tasks, err := services.GetTasksByGoalID(uint(goalID))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Не удалось получить задачи")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDParam(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Некорректный ID")
		return
	}

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный JSON")
		return
	}

	task, err := services.UpdateTask(uint(id), updatedTask)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDParam(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Некорректный ID")
		return
	}

	if err := services.DeleteTask(uint(id)); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Не удалось удалить задачу")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Задача удалена"})
}

func GetTeamTasks(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Не авторизован")
		return
	}

	tasks, err := services.GetTeamTasks(user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Не удалось получить командные задачи")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tasks)
}

func GetAllTeamTasks(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil || user.Role != "superadmin" {
		utils.RespondWithError(w, http.StatusForbidden, "Доступ только для суперадмина")
		return
	}

	tasks, err := services.GetAllTeamTasks()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка получения задач")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tasks)
}

func GetAllTeamTasksKanban(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil || user.Role != "superadmin" {
		utils.RespondWithError(w, http.StatusForbidden, "Доступ только для суперадмина")
		return
	}

	tasks, err := services.GetAllTeamTasks()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка получения задач")
		return
	}

	result := make(map[string]map[string][]models.Task)

	for _, task := range tasks {
		if task.Team == nil {
			continue // если нет команды — пропускаем
		}
		teamName := task.Team.Name
		if _, ok := result[teamName]; !ok {
			result[teamName] = map[string][]models.Task{
				"todo": {},
				"done": {},
			}
		}

		if task.IsDone {
			result[teamName]["done"] = append(result[teamName]["done"], task)
		} else {
			result[teamName]["todo"] = append(result[teamName]["todo"], task)
		}
	}

	utils.RespondWithJSON(w, http.StatusOK, result)
}

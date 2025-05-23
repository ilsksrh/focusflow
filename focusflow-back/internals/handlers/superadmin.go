package handlers

import (
	"encoding/json"
	"focusflow/internals/services"
	"focusflow/internals/utils"
	"net/http"
	"strconv"
)

func isSuperadmin(r *http.Request) bool {
	user, _ := utils.GetCurrentUser(r)
	return user.Role == "superadmin"
}

func SuperadminGetTeams(w http.ResponseWriter, r *http.Request) {
	if !isSuperadmin(r) {
		utils.RespondWithError(w, http.StatusForbidden, "Только для суперадмина")
		return
	}

	teams, err := services.GetAllTeams()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка получения команд")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, teams)
}

func SuperadminCreateTeam(w http.ResponseWriter, r *http.Request) {
	if !isSuperadmin(r) {
		utils.RespondWithError(w, http.StatusForbidden, "Только для суперадмина")
		return
	}

	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный JSON")
		return
	}

	if err := services.CreateTeam(input.Name); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Команда создана"})
}

func SuperadminGetTeamByID(w http.ResponseWriter, r *http.Request) {
	if !isSuperadmin(r) {
		utils.RespondWithError(w, http.StatusForbidden, "Только для суперадмина")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "Некорректный ID")
		return
	}

	team, err := services.GetTeamByID(uint(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Команда не найдена")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, team)
}

func SuperadminDeleteTeam(w http.ResponseWriter, r *http.Request) {
	if !isSuperadmin(r) {
		utils.RespondWithError(w, http.StatusForbidden, "Только для суперадмина")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "Некорректный ID")
		return
	}

	if err := services.DeleteTeam(uint(id)); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка удаления команды")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Команда удалена"})
}

func SuperadminAssignUserToTeam(w http.ResponseWriter, r *http.Request) {
	if !isSuperadmin(r) {
		utils.RespondWithError(w, http.StatusForbidden, "Только для суперадмина")
		return
	}

	var input struct {
		UserID uint `json:"user_id"`
		TeamID uint `json:"team_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный JSON")
		return
	}

	if err := services.AssignUserToTeam(input.UserID, input.TeamID); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка назначения пользователя")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Пользователь назначен в команду"})
}

func SuperadminRemoveUserFromTeam(w http.ResponseWriter, r *http.Request) {
	if !isSuperadmin(r) {
		utils.RespondWithError(w, http.StatusForbidden, "Только для суперадмина")
		return
	}

	var input struct {
		UserID uint `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Неверный JSON")
		return
	}

	if err := services.RemoveUserFromTeam(input.UserID); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка удаления пользователя из команды")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Пользователь удалён из команды"})
}

func GetMyTeam(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Не авторизован")
		return
	}

	if user.TeamID == nil {
		utils.RespondWithError(w, http.StatusNotFound, "Вы не состоите в команде")
		return
	}

	team, err := services.GetTeamByID(*user.TeamID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Ошибка загрузки команды")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, team)
}


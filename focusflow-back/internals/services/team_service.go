package services

import (
	"errors"
	"focusflow/internals/db"
	"focusflow/internals/models"
)

func CreateTeam(name string) error {
	if name == "" {
		return errors.New("Название команды не может быть пустым")
	}
	return db.DB.Create(&models.Team{Name: name}).Error
}

func DeleteTeam(id uint) error {

	if err := db.DB.Model(&models.User{}).Where("team_id = ?", id).Update("team_id", nil).Error; err != nil {
		return err
	}


	return db.DB.Delete(&models.Team{}, id).Error
}

func GetTeamByID(id uint) (models.Team, error) {
	var team models.Team
	if err := db.DB.Preload("Users").First(&team, id).Error; err != nil {
		return models.Team{}, err
	}
	return team, nil
}


func GetAllTeams() ([]models.Team, error) {
	var teams []models.Team
	if err := db.DB.Preload("Users").Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}


func AssignUserToTeam(userID, teamID uint) error {
	// Проверка существования команды
	var team models.Team
	if err := db.DB.First(&team, teamID).Error; err != nil {
		return errors.New("Команда не найдена")
	}

	// Проверка существования пользователя
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return errors.New("Пользователь не найден")
	}

	// Назначение
	return db.DB.Model(&user).Update("team_id", teamID).Error
}

func RemoveUserFromTeam(userID uint) error {
	return db.DB.Model(&models.User{}).Where("id = ?", userID).Update("team_id", nil).Error
}


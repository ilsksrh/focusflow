package services

import (
	"errors"
	"focusflow/internals/db"
	"focusflow/internals/models"
	"fmt"
)

// Создать новую цель
func CreateGoal(goal *models.Goal) error {
	return db.DB.Create(goal).Error
}

// Получить все цели
func GetGoals() ([]models.Goal, error) {
	var goals []models.Goal

	if err := db.DB.Preload("Tasks.Tags").Find(&goals).Error; err != nil {
		return nil, err
	}

	// Считаем прогресс каждой цели
	for i := range goals {
		total := len(goals[i].Tasks)
		done := 0
		for _, task := range goals[i].Tasks {
			if task.IsDone {
				done++
			}
		}
		if total == 0 {
			goals[i].Progress = "0%"
		} else {
			percent := (float64(done) / float64(total)) * 100
			goals[i].Progress = fmt.Sprintf("%.0f%%", percent)
		}
	}

	return goals, nil
}

// Обновить цель по ID
func UpdateGoal(id uint, updatedGoal models.Goal) (models.Goal, error) {
	var goal models.Goal
	if err := db.DB.First(&goal, id).Error; err != nil {
		return models.Goal{}, errors.New("Цель не найдена")
	}

	goal.Title = updatedGoal.Title
	goal.Description = updatedGoal.Description

	if err := db.DB.Save(&goal).Error; err != nil {
		return models.Goal{}, err
	}

	return goal, nil
}

// Удалить цель по ID
func DeleteGoal(id uint) error {
	return db.DB.Delete(&models.Goal{}, id).Error
}

func GetGoalsByUser(userID uint) ([]models.Goal, error) {
    var goals []models.Goal
    err := db.DB.Where("user_id = ?", userID).Preload("Tasks.Tags").Find(&goals).Error
    return goals, err
}

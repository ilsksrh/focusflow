package services

import (
	"errors"
	"focusflow/internals/db"
	"focusflow/internals/models"
)

// Создание задачи с тегами
func CreateTask(task *models.Task) error {
	task.ID = 0

	var processedTags []*models.Tag
	for _, tag := range task.Tags {
		var existing models.Tag
		if err := db.DB.Where("name = ?", tag.Name).First(&existing).Error; err == nil {
			processedTags = append(processedTags, &existing)
		} else {
			newTag := models.Tag{Name: tag.Name}
			if newTag.Name != "" {
				db.DB.Create(&newTag)
				processedTags = append(processedTags, &newTag)
			}
		}
	}

	task.Tags = processedTags

	return db.DB.Create(task).Error
}

// Получение всех задач
func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := db.DB.Preload("Tags").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// Получение задачи по ID
func GetTaskByID(id uint) (models.Task, error) {
	var task models.Task
	if err := db.DB.Preload("Tags").First(&task, id).Error; err != nil {
		return models.Task{}, errors.New("Задача не найдена")
	}
	return task, nil
}

// Получение всех задач по ID цели
func GetTasksByGoalID(goalID uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := db.DB.Where("goal_id = ?", goalID).Preload("Tags").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// Обновление задачи
func UpdateTask(id uint, updatedTask models.Task) (models.Task, error) {
	var task models.Task
	if err := db.DB.First(&task, id).Error; err != nil {
		return models.Task{}, errors.New("Задача не найдена")
	}

	task.Title = updatedTask.Title
	task.IsDone = updatedTask.IsDone
	task.Description = updatedTask.Description

	if err := db.DB.Save(&task).Error; err != nil {
		return models.Task{}, err
	}

	return task, nil
}


// Удаление задачи
func DeleteTask(id uint) error {
	return db.DB.Delete(&models.Task{}, id).Error
}

func GetTasksByUser(userID uint) ([]models.Task, error) {
    var tasks []models.Task
    err := db.DB.Where("user_id = ?", userID).Preload("Tags").Find(&tasks).Error
    return tasks, err
}

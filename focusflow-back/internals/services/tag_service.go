package services

import (
	"errors"
	"focusflow/internals/db"
	"focusflow/internals/models"
)

// Создать тег (если его ещё нет)
func CreateTag(tag *models.Tag) (*models.Tag, error) {
	var existing models.Tag
	if err := db.DB.Where("LOWER(name) = LOWER(?)", tag.Name).First(&existing).Error; err == nil {
		return &existing, nil
	}

	if err := db.DB.Create(tag).Error; err != nil {
		return nil, err
	}

	return tag, nil
}

// Получить все теги
func GetTags() ([]models.Tag, error) {
	var tags []models.Tag
	if err := db.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

// Получить теги по ID задачи
func GetTagsByTaskID(taskID uint) ([]*models.Tag, error) {
	var task models.Task
	if err := db.DB.Preload("Tags").First(&task, taskID).Error; err != nil {
		return nil, errors.New("Задача не найдена")
	}

	return task.Tags, nil
}

// Удалить тег
func DeleteTag(id uint) error {
	return db.DB.Delete(&models.Tag{}, id).Error
}

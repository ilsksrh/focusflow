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
		if tag.ID == 0 {
			continue
		}
		var existing models.Tag
		if err := db.DB.First(&existing, tag.ID).Error; err == nil {
			processedTags = append(processedTags, &existing)
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

func UpdateTask(id uint, updatedTask models.Task) (models.Task, error) {
	var task models.Task
	if err := db.DB.Preload("Tags").First(&task, id).Error; err != nil {
		return models.Task{}, errors.New("Задача не найдена")
	}

	// обновление основных полей
	task.Title = updatedTask.Title
	task.IsDone = updatedTask.IsDone
	task.Description = updatedTask.Description
	task.GoalID = updatedTask.GoalID

	// загрузка тегов
	var newTags []*models.Tag
	for _, tag := range updatedTask.Tags {
		if tag.ID == 0 {
			continue
		}
		var existing models.Tag
		if err := db.DB.First(&existing, tag.ID).Error; err == nil {
			newTags = append(newTags, &existing)
		}
	}

	// обновление связи many2many
	if err := db.DB.Model(&task).Association("Tags").Replace(newTags); err != nil {
		return models.Task{}, err
	}

	if err := db.DB.Save(&task).Error; err != nil {
		return models.Task{}, err
	}

	// загрузка тегов заново, чтобы вернуть в ответе
	db.DB.Preload("Tags").First(&task, id)
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

func GetTeamTasks(userID uint) ([]models.Task, error) {
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	var tasks []models.Task
	err := db.DB.
		Joins("JOIN task_tags ON task_tags.task_id = tasks.id").
		Joins("JOIN tags ON tags.id = task_tags.tag_id").
		Where("tasks.team_id = ? AND LOWER(tags.name) = ?", user.TeamID, "командная").
		Preload("Tags").
		Preload("User").
		Preload("Team").
		Find(&tasks).Error

	return tasks, err
}

func GetAllTeamTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := db.DB.Joins("JOIN task_tags ON task_tags.task_id = tasks.id").
		Joins("JOIN tags ON tags.id = task_tags.tag_id").
		Where("LOWER(tags.name) = ?", "командная").
		Preload("Tags").
		Preload("User").
		Preload("Team").
		Find(&tasks).Error

	return tasks, err
}

func GetKanbanByTeams() (map[string]map[string][]models.Task, error) {
	var teams []models.Team
	if err := db.DB.Preload("Users").Find(&teams).Error; err != nil {
		return nil, err
	}

	result := make(map[string]map[string][]models.Task)

	for _, team := range teams {
		var tasks []models.Task
		err := db.DB.
			Joins("JOIN task_tags ON task_tags.task_id = tasks.id").
			Joins("JOIN tags ON tags.id = task_tags.tag_id").
			Where("tasks.team_id = ? AND LOWER(tags.name) = ?", team.ID, "командная").
			Preload("Tags").
			Preload("User").
			Find(&tasks).Error

		if err != nil {
			continue 

		todo := []models.Task{}
		done := []models.Task{}

		for _, t := range tasks {
			if t.IsDone {
				done = append(done, t)
			} else {
				todo = append(todo, t)
			}
		}

		result[team.Name] = map[string][]models.Task{
			"todo": todo,
			"done": done,
		}
	}
}

	return result, nil
}

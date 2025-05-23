package models

import "gorm.io/gorm"

type Task struct {
    gorm.Model
    Title       string    `json:"title"`
    IsDone      bool      `json:"is_done"`
    GoalID      uint      `json:"goal_id"`
    Description string    `json:"description"`
    Tags        []*Tag    `gorm:"many2many:task_tags;" json:"tags"`
    UserID      uint      `json:"user_id"`


}

package models

import "gorm.io/gorm"

type Tag struct {
    gorm.Model
    Name  string   `json:"name"`
    Tasks []Task   `gorm:"many2many:task_tags;" json:"-"`
}

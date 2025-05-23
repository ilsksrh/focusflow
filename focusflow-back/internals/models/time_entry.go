package models

import "gorm.io/gorm"

type TimeEntry struct {
    gorm.Model
    TaskID         uint   `json:"task_id"`
    DurationMinute int    `json:"duration_minute"`
    Note           string `json:"note"`
    Date           string `json:"date"`
}

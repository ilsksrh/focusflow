package models

type Team struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"unique" json:"name"`
	Users []User
}

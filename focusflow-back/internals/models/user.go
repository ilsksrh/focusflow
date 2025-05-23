package models


type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"unique" json:"username"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
	IsBlocked    bool   `gorm:"default:false" json:"is_blocked"`
	ProfilePicture string `json:"profile_picture"`
}

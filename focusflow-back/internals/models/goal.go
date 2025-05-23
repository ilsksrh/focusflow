package models


type Goal struct {
    ID          uint   `json:"id"` 
    Title       string `json:"title"`
    Description string `json:"description"`
    Progress    string `json:"progress" gorm:"-"`
    Tasks       []Task `json:"tasks" gorm:"foreignKey:GoalID"`
    UserID      uint   `json:"user_id"`
    Date        string `json:"date"`
}

package db

import (
    "focusflow/internals/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func Init() {
    dsn := "host=localhost user=postgres password=sara1234 dbname=focusflow port=5432 sslmode=disable"
    // log.Println("DSN:", dsn)


    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    err = db.AutoMigrate(&models.User{}, &models.Goal{}, &models.Task{}, &models.Tag{}, &models.TimeEntry{})

    
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    DB = db
}

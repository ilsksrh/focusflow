package main

import (
    "focusflow/configs"
    "focusflow/internals/db"
    "focusflow/internals/models"
    "focusflow/internals/routes"
    "focusflow/internals/middlewares"
    "log"
    "net/http"
)

func main() {
	config := configs.NewConfig()
	db.Init()

	err := db.DB.AutoMigrate(&models.Goal{}, &models.Task{}, &models.Tag{}, &models.TimeEntry{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate DB:", err)
	}

	router := routes.RegisterRoutes()

	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

    http.Handle("/", middlewares.CORS(router))


	log.Printf("Server is running at http://%s\n", config.ServerAddress)
	log.Fatal(http.ListenAndServe(config.ServerAddress, nil))
}

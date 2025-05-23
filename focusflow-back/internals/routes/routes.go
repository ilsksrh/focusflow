package routes

import (
	"focusflow/internals/handlers"
	"focusflow/internals/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middlewares.CORS)

	r.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	// Главная страница
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FocusFlow API is running!"))
	})

	// --- Аутентификация ---
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/logout", handlers.Logout).Methods("POST")
	r.Handle("/me", middlewares.RequireAuth(http.HandlerFunc(handlers.Me))).Methods("GET")

	// --- Цели (Goals) ---
	goalsRouter := r.PathPrefix("/goals").Subrouter()
	goalsRouter.Use(middlewares.RequireAuth)
	goalsRouter.HandleFunc("", handlers.GetGoals).Methods("GET")
	goalsRouter.HandleFunc("", handlers.CreateGoal).Methods("POST")
	goalsRouter.HandleFunc("/{id}", handlers.GetTasksByGoalID).Methods("GET")
	goalsRouter.HandleFunc("/{id}", handlers.UpdateGoal).Methods("PUT")
	goalsRouter.HandleFunc("/{id}", handlers.DeleteGoal).Methods("DELETE")

	// --- Задачи (Tasks) ---
	tasksRouter := r.PathPrefix("/tasks").Subrouter()
	tasksRouter.Use(middlewares.RequireAuth)
	tasksRouter.HandleFunc("", handlers.GetTasks).Methods("GET")
	tasksRouter.HandleFunc("", handlers.CreateTask).Methods("POST")
	tasksRouter.HandleFunc("/{id}", handlers.GetTaskByID).Methods("GET")
	tasksRouter.HandleFunc("/{id}", handlers.UpdateTask).Methods("PUT")
	tasksRouter.HandleFunc("/{id}", handlers.DeleteTask).Methods("DELETE")
	tasksRouter.HandleFunc("/{id}/tags", handlers.GetTagsByTaskID).Methods("GET")

	// --- Теги (Tags) ---
	tagsRouter := r.PathPrefix("/tags").Subrouter()
	tagsRouter.Use(middlewares.RequireAuth)
	tagsRouter.HandleFunc("", handlers.GetTags).Methods("GET")
	tagsRouter.HandleFunc("", handlers.CreateTag).Methods("POST")
	tagsRouter.HandleFunc("/{id}", handlers.DeleteTag).Methods("DELETE")

	// --- Админ-панель ---
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middlewares.RequireAdmin)
	adminRouter.HandleFunc("/users", handlers.AdminGetUsers).Methods("GET")
	adminRouter.HandleFunc("/add_user", handlers.AdminAddUser).Methods("POST")
	adminRouter.HandleFunc("/block_user", handlers.AdminBlockUser).Methods("POST")
	adminRouter.HandleFunc("/unblock_user", handlers.AdminUnblockUser).Methods("POST")
	adminRouter.HandleFunc("/delete_user", handlers.AdminDeleteUser).Methods("DELETE")

	// --- Работа с профилем ---
	profile := r.PathPrefix("/profile").Subrouter()
	profile.Use(middlewares.RequireAuth)
	profile.HandleFunc("/upload_picture", handlers.UploadProfilePicture).Methods("POST")
	profile.HandleFunc("/change_password", handlers.ChangePassword).Methods("PUT")
	profile.HandleFunc("/update", handlers.UpdateProfile).Methods("PUT")

	return r
}

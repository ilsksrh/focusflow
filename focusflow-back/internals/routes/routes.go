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
	goalsRouter.HandleFunc("/{id}/tasks", handlers.GetTasksByGoalID).Methods("GET")
	goalsRouter.HandleFunc("/{id}", handlers.UpdateGoal).Methods("PUT")
	goalsRouter.HandleFunc("/{id}", handlers.DeleteGoal).Methods("DELETE")

	// --- Задачи (Tasks) ---
	tasksRouter := r.PathPrefix("/tasks").Subrouter()
	tasksRouter.Use(middlewares.RequireAuth)
	tasksRouter.HandleFunc("", handlers.GetTasks).Methods("GET")
	tasksRouter.HandleFunc("", handlers.CreateTask).Methods("POST")
	tasksRouter.HandleFunc("/team", handlers.GetTeamTasks).Methods("GET")
	tasksRouter.HandleFunc("/{id}", handlers.GetTaskByID).Methods("GET")
	tasksRouter.HandleFunc("/{id}", handlers.UpdateTask).Methods("PUT")
	tasksRouter.HandleFunc("/{id}", handlers.DeleteTask).Methods("DELETE")
	// tasksRouter.HandleFunc("/all_teams", handlers.GetAllTeamTasks).Methods("GET")


	tasksRouter.HandleFunc("/{id}/tags", handlers.GetTagsByTaskID).Methods("GET")
	
	


	// --- Теги (Tags) ---
	tagsRouter := r.PathPrefix("/tags").Subrouter()
	tagsRouter.Use(middlewares.RequireAuth)
	tagsRouter.HandleFunc("", handlers.GetTags).Methods("GET")
	tagsRouter.HandleFunc("", handlers.CreateTag).Methods("POST")
	tagsRouter.HandleFunc("/{id}", handlers.DeleteTag).Methods("DELETE")
	tagsRouter.HandleFunc("/{id}", handlers.UpdateTag).Methods("PUT")


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

	// --- Суперадмин: команды (Teams) ---
	superadmin := r.PathPrefix("/superadmin").Subrouter()
	superadmin.Use(middlewares.RequireSuperadmin) 
	superadmin.HandleFunc("/teams", handlers.SuperadminGetTeams).Methods("GET")
	superadmin.HandleFunc("/teams", handlers.SuperadminCreateTeam).Methods("POST")
	superadmin.HandleFunc("/teams", handlers.SuperadminDeleteTeam).Methods("DELETE") 
	superadmin.HandleFunc("/assign", handlers.SuperadminAssignUserToTeam).Methods("POST")
	superadmin.HandleFunc("/remove_from_team", handlers.SuperadminRemoveUserFromTeam).Methods("POST")
	superadmin.HandleFunc("/team", handlers.SuperadminGetTeamByID).Methods("GET")
	

	r.Handle("/my_team", middlewares.RequireAuth(http.HandlerFunc(handlers.GetMyTeam))).Methods("GET")
	tasksRouter.HandleFunc("/kanban/all_teams", handlers.GetAllTeamTasksKanban).Methods("GET")





	return r
}

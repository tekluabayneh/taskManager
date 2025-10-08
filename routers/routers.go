package routers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/tekluabayney/taskmanger/handlers"

	db "github.com/tekluabayney/taskmanger/internal/db"
)

func LoadRouter(db *db.Queries) *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"opps! it's a live"}`))

	})

	router.Route("/tasks", func(r chi.Router) {
		TaskRouter(r, db)
	},
	)
	router.Route("/users", func(r chi.Router) {
		UserRouter(r, db)
	},
	)

	return router
}

func TaskRouter(router chi.Router, q *db.Queries) {
	TaskHandeler := &handlers.Taskhandler{
		DB: q,
	}
	router.Get("/getTask", TaskHandeler.GetTask)
	router.Get("/updateTask", TaskHandeler.UpdateTask)
	router.Get("/deleteTask", TaskHandeler.DeleteTask)
	router.Get("/getSingleTask", TaskHandeler.GetSingleTask)

}

func UserRouter(router chi.Router, q *db.Queries) {
	UserHnadler := &handlers.UserType{
		DB: q,
	}

	router.Get("/getUser", UserHnadler.GetUser)
	router.Get("/updateUser", UserHnadler.UpdateUser)
	router.Get("/deleteUser", UserHnadler.DeleteUser)
	router.Get("/getSingleUser", UserHnadler.InsertUser)

}

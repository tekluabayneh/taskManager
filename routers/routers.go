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

	router.Route("/getTask", func(r chi.Router) {
		getTasks(r, db)
	},
	)
	router.Route("/updateTask", func(r chi.Router) {
		getTasks(r, db)
	},
	)

	return router
}

func getTasks(router chi.Router, q *db.Queries) {
	getTaskRoute := &handlers.Taskhandler{
		DB: q,
	}

	router.Get("/", getTaskRoute.GetTask)
}

func UpdateTask(router chi.Router, q *db.Queries) {
	updateTasksHandler := &handlers.Taskhandler{
		DB: q,
	}
	router.Get("/", updateTasksHandler.UpdateTask)

}

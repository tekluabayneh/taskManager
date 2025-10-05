package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"net/http"

	"github.com/tekluabayney/taskmanger/handlers"
)

func LoadRouter() *chi.Mux {
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
		w.Write([]byte(`{"message":"it works man"}`))

	})

	router.Route("/getTask", getTasks)
	router.Route("/updatTask", getTasks)

	return router
}

func getTasks(router chi.Router) {
	getTaskRoute := &handlers.Task{}
	router.Get("/", getTaskRoute.GetTask)
}

func UpdateTask(router chi.Router) {
	updateTasksHandler := &handlers.Task{}
	router.Put("/", updateTasksHandler.updatTask)

}

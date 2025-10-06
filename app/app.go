package app

import (
	"fmt"
	"github.com/tekluabayney/taskmanger/routers"
	"log"
	"net/http"
	"os"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: routers.LoadRouter(),
	}

	return app
}

func (app *App) Start() {
	PORT := os.Getenv("PORT")
	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: app.router,
	}

	fmt.Println("Server is running at http://localhost:3000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to serve server %v", err)

	}

}

package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/tekluabayney/taskmanger/internal/db"
	"github.com/tekluabayney/taskmanger/routers"
)

type App struct {
	router http.Handler
	db     *db.Queries
	pool   *pgxpool.Pool
}

func New() *App {
	DBURL := os.Getenv("DB_URL")
	if DBURL == "" {
		log.Fatal("DB_URL is empty")
	}

	pool, err := pgxpool.New(context.Background(), DBURL)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v\n", err)
	}

	// Create SQLC queries instance
	db := db.New(pool)

	app := &App{
		db:     db,
		pool:   pool,
		router: routers.LoadRouter(db),
	}
	return app
}

func (app *App) Start() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9090"
	}

	defer app.pool.Close()

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: app.router,
	}

	fmt.Printf("Server running at http://localhost:%s\n", PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

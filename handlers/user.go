package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	db "github.com/tekluabayney/taskmanger/internal/db"
)

type UserType struct {
	DB *db.Queries
}

func (h *UserType) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.DB.GetsingleUser(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}

func (h *UserType) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *UserType) InsertUser(q http.ResponseWriter, r *http.Request) {

}
func (h *UserType) DeleteUser(q http.ResponseWriter, r *http.Request) {

}

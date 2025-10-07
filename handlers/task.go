package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	db "github.com/tekluabayney/taskmanger/internal/db"
)

type Taskhandler struct {
	DB *db.Queries
}

func (h *Taskhandler) GetTask(w http.ResponseWriter, r *http.Request) {
	user, err := h.DB.GetUser(context.Background(), 2)
	if err != nil {
		fmt.Println("failed toget user")
	}
	fmt.Println("get task")
	fmt.Println(user)
}

func (h *Taskhandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("udpate task")
	type newUser struct {
		name string
		age  int
	}
	nUser := db.CreateUserParams{
		Name:  "one",
		Email: "man@gmail.com",
	}
	user, err := h.DB.CreateUser(context.Background(), nUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

}

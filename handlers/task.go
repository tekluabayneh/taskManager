package handlers

import (
	_ "context"
	_ "encoding/json"
	"fmt"
	"net/http"

	db "github.com/tekluabayney/taskmanger/internal/db"
)

type Taskhandler struct {
	DB *db.Queries
}
type taskFormater struct {
	Id          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      string `json:"userid"`
}

func (h *Taskhandler) GetTask(w http.ResponseWriter, r *http.Request) {
	task, err := h.DB.GetTasks(context.Background())
	if err != nil {
		http.Error(w, "failed toget user", http.StatusInternalServerError)
		return
	}

	result := formatType{
		Id:    task
		Name:  user.Name,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}
func (h *Taskhandler) NewTask(w http.ResponseWriter, r *http.Request) {
	var inputFormater taskFormater

	err := json.NewDecoder(r.Body).Decode(&inputFormater)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	task, err := h.DB.CreateTask(context.Background(), inputFormater)

	if err != nil {
		http.Error(w, "task didn't created", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}

func (h *Taskhandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	type newUser struct {
		name string
		age  int
	}

	nUser := db.CreateUserParams{
		Name:  "one",
		Email: "man@gmail.com",
	}

	user, err := h.DB.(context.Background(), nUser)
	if err != nil {
		log.Fatal(err)
	}

}

func (h *Taskhandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("task delted")
}

func (h *Taskhandler) GetSingleTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get single task")
}

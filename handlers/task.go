package handlers

import (
	"fmt"
	"net/http"
)

type Task struct{}

func (T *Task) GetTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get task")
}

func (T *Task) UpdatTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("udpate task")
}

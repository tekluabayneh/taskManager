package handlers

import (
	"fmt"
	"net/http"
)

type Taskhandler struct{}

func (T *Taskhandler) GetTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println("get task")
}

func (T *Taskhandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("udpate task")
}

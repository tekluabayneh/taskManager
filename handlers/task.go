package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/tekluabayney/taskmanger/internal/db"
)

type Taskhandler struct {
	DB *db.Queries
}
type taskFormater struct {
	Id          int32       `json:"id"`
	Title       string      `json:"title"`
	Status      pgtype.Text `json:"status"`
	Description pgtype.Text `json:"description"`
	UserId      pgtype.Int4 `json:"userid"`
}

func (h *Taskhandler) GetTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var inputType taskFormater
	json.NewDecoder(r.Body).Decode(&inputType)

	tasks, err := h.DB.GetTasks(context.Background())

	if err != nil {
		http.Error(w, "failed toget user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)

}

func (h *Taskhandler) NewTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inputFormater taskFormater

	err := json.NewDecoder(r.Body).Decode(&inputFormater)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	task := db.CreateTaskParams{
		Title:       inputFormater.Title,
		Description: inputFormater.Description,
		Status:      inputFormater.Status,
		UserID:      inputFormater.UserId,
	}

	Rtask, err := h.DB.CreateTask(context.Background(), task)

	if err != nil {
		http.Error(w, "task didn't created", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(Rtask)

}

func (h *Taskhandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inputType taskFormater

	err := json.NewDecoder(r.Body).Decode(&inputType)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid json",
		})
	}

	upTask := db.UpdateTaskParams{
		ID:          inputType.Id,
		Title:       inputType.Title,
		Status:      inputType.Status,
		Description: inputType.Description,
	}

	newTask, err := h.DB.UpdateTask(context.Background(), upTask)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid json",
		})

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"msg": "task updated", "task": newTask,
	})

}

func (h *Taskhandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	strId := r.URL.Query().Get("id")

	if strId == "" {
		http.Error(w, `{"error":"missing id "}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(strId)

	if err != nil {
		http.Error(w, `{"error":"invalid id "}`, http.StatusBadRequest)
		return
	}
	_, err = h.DB.GetSingTask(context.Background(), int32(id))
	if err != nil {
		http.Error(w, `{"error": "task not found"}`, http.StatusNotFound)
		return
	}

	DeletedTask, err := h.DB.DeleteTask(context.Background(), int32(id))
	if err != nil {
		http.Error(w, `{"error": "failed to delete task"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message":    "task deleted successfully",
		"deltedTask": DeletedTask,
	})

}

func (h *Taskhandler) GetSingleTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, `{"error": "missing id parameter"}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, `{"error": "invalid id"}`, http.StatusBadRequest)
		return
	}

	task, err := h.DB.GetSingTask(context.Background(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, `{"error": "task not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error": "failed to get task"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "task retrieved successfully",
		"task":    task,
	})

}

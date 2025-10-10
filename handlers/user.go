package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	db "github.com/tekluabayney/taskmanger/internal/db"
)

type UserType struct {
	DB *db.Queries
}

type formatType struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserType) GetUsers(w http.ResponseWriter, r *http.Request) {
	user, err := h.DB.GetUser(context.Background())
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	result := formatType{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}

func (h *UserType) GetSingUser(w http.ResponseWriter, r *http.Request) {

	var inputtpe formatType
	err := json.NewDecoder(r.Body).Decode(&inputtpe)
	user, err := h.DB.GetsingleUser(context.Background(), inputtpe.Id)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	result := formatType{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *UserType) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var newtype formatType
	err := json.NewDecoder(r.Body).Decode(&newtype)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid json",
		})
		http.Error(w, "User not updated", http.StatusBadRequest)
		return
	}

	data := db.UpdateUserParams{
		ID:    newtype.Id,
		Name:  newtype.Name,
		Email: newtype.Email,
	}

	NewUser, err := h.DB.UpdateUser(context.Background(), data)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid json",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"msg": NewUser,
	})

}

func (h *UserType) InsertUser(w http.ResponseWriter, r *http.Request) {
	var inputtpe formatType

	err := json.NewDecoder(r.Body).Decode(&inputtpe)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid json",
		})
		return

	}

	rData := db.CreateUserParams{
		Name:  inputtpe.Name,
		Email: inputtpe.Email,
	}

	data, err := h.DB.CreateUser(context.Background(), rData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"erro": "internal server error ", "msg": err})
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)

}

func (h *UserType) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var inputtpe formatType
	json.NewDecoder(r.Body).Decode(&inputtpe)
	user, err := h.DB.DeleteUser(context.Background(), inputtpe.Id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"erro": "internal server error ", "msg": err})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"msg": "user is deleted ", "deleted_user": user})

}

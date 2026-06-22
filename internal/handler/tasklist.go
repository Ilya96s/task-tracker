package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/task-tracker/internal/service"
)

type CreateTaskListRequest struct {
	Name string
}

type TaskListHandler struct {
	service *service.TaskListService
}

func NewTaskListHandler(service *service.TaskListService) *TaskListHandler {
	return &TaskListHandler{
		service: service,
	}
}

func (h *TaskListHandler) CreateList(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskListRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	list, err := h.service.Create(r.Context(), req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TaskListHandler) GetList(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("list_id")

	if len(pathValue) == 0 {
		http.Error(w, "list_id query parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	taskList, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(taskList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

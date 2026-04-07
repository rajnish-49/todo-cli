package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"todo-cli/internal/todo"
)

type createTaskRequest struct {
	Title string `json:"title"`
}

type Handler struct {
	service *todo.Service
}

func NewHandler(service *todo.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// r is what the client sent
// w is what the server will send back
func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := h.service.ListTasks()
	if err != nil {
		http.Error(w, "failed to load tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func parseTaskID(path string) (int, error) {
	trimmedPath := strings.TrimPrefix(path, "/tasks/")
	parts := strings.Split(trimmedPath, "/")

	return strconv.Atoi(parts[0])
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	task, err := h.service.AddTask(req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) MarkDone(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasSuffix(r.URL.Path, "/done") {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	id, err := parseTaskID(r.URL.Path)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.service.MarkDone(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := parseTaskID(r.URL.Path)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

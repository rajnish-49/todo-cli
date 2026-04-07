package main

import (
	"fmt"
	"net/http"
	"os"

	"todo-cli/cmd"
	"todo-cli/internal/api"
	"todo-cli/internal/storage"
	"todo-cli/internal/todo"
)

func main() {
	store := storage.NewFileStore("tasks.json")
	service := todo.NewService(store)
	handler := cmd.NewCommandHandler(service)
	apiHandler := api.NewHandler(service)

	if len(os.Args) > 1 && os.Args[1] == "serve" {
		http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				apiHandler.ListTasks(w, r)
				fmt.Println("successfull request ")
			case http.MethodPost:
				apiHandler.CreateTask(w, r)
			default:
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			}
		})

		http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPatch:
				apiHandler.MarkDone(w, r)
			case http.MethodDelete:
				apiHandler.DeleteTask(w, r)
			default:
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			}
		})

		fmt.Println("Server running on http://localhost:8080")

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		return
	}

	err := handler.Execute(os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

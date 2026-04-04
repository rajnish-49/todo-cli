package main

import (
	"fmt"
	"os"

	"todo-cli/cmd"
	"todo-cli/internal/storage"
	"todo-cli/internal/todo"
)

func main() {
	store := storage.NewFileStore("tasks.json")
	service := todo.NewService(store)
	handler := cmd.NewCommandHandler(service)

	err := handler.Execute(os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}


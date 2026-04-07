package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"todo-cli/internal/todo"
)

// service is stored inside that struct
type CommandHandler struct {
	service *todo.Service
}

func NewCommandHandler(service *todo.Service) *CommandHandler {
	return &CommandHandler{
		service: service,
	}
}

func (h *CommandHandler) Execute(args []string) error {
	if len(args) == 0 {
		h.Help()
		return nil
	}

	switch args[0] {
	case "list":
		return h.List()
	case "add":
		if len(args) < 2 {
			return errors.New("please provide a task title")
		}

		title := strings.Join(args[1:], " ")
		return h.Add(title)
	case "done":
		if len(args) < 2 {
			return errors.New("please provide a task id")
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid task id: %w", err)
		}

		return h.Done(id)
	case "delete":
		if len(args) < 2 {
			return errors.New("please provide a task id")
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid task id: %w", err)
		}

		return h.Delete(id)
	default:
		h.Help()
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func (h *CommandHandler) Help() {
	fmt.Println("Usage:")
	fmt.Println("  todo list")
	fmt.Println("  todo add <title>")
	fmt.Println("  todo done <id>")
	fmt.Println("  todo delete <id>")
}

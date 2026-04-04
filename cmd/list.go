package cmd

import "fmt"

func (h *CommandHandler) List() error {
	tasks, err := h.service.ListTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}

		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
	}

	return nil
}

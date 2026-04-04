package cmd

import "fmt"

func (h *CommandHandler) Add(title string) error {
	task, err := h.service.AddTask(title)
	if err != nil {
		return err
	}

	fmt.Printf("Added task %d: %s\n", task.ID, task.Title)
	return nil
}

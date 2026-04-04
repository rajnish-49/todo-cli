package cmd

import "fmt"

func (h *CommandHandler) Delete(id int) error {
	err := h.service.DeleteTask(id)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted task %d.\n", id)
	return nil
}

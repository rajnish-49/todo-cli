package cmd

import "fmt"

func (h *CommandHandler) Done(id int) error {
	err := h.service.MarkDone(id)
	if err != nil {
		return err
	}

	fmt.Printf("Marked task %d as done.\n", id)
	return nil
}

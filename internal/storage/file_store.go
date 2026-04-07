package storage

import (
	"encoding/json"
	"os"

	"todo-cli/internal/todo"
)

type FileStore struct {
	FilePath string
}

func NewFileStore(filePath string) *FileStore {
	return &FileStore{
		FilePath: filePath,
	}
}

func (fs *FileStore) Load() ([]todo.Task, error) {
	data, err := os.ReadFile(fs.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []todo.Task{}, nil
		}

		return nil, err
	}

	var tasks []todo.Task

	if len(data) == 0 {
		return []todo.Task{}, nil
	}

	// converts json into Go data structure (slice of Task)
	err = json.Unmarshal(data, &tasks) // returns either error or nil
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (fs *FileStore) Save(tasks []todo.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fs.FilePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

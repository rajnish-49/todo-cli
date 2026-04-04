package todo

import (
	"errors"
	"strings"
	"time"
)

type Store interface {
	Load() ([]Task, error)
	Save([]Task) error
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

// method that belongs tp Service struct 
func (s *Service) ListTasks() ([]Task, error) {
	tasks, err := s.store.Load()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Service) AddTask(title string) (Task, error) {
	tasks, err := s.store.Load()
	if err != nil {
		return Task{}, err
	}

	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, errors.New("title can't be empty")
	}

	newTask := Task{
		ID:        nextTaskID(tasks),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)
	err = s.store.Save(tasks)
	if err != nil {
		return Task{}, err
	}

	return newTask, nil
}

func (s *Service) MarkDone(id int) error {
	tasks, err := s.store.Load()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			return s.store.Save(tasks)
		}
	}

	return errors.New("task not found")
}

func nextTaskID(tasks []Task) int {
	maxID := 0

	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}

func (s *Service) DeleteTask(id int) error {
	tasks, err := s.store.Load()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return s.store.Save(tasks)
		}
	}

	return errors.New("task not found")
}


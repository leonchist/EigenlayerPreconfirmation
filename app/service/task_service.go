package service

import (
	"context"

	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/repository"
)

type TaskServiceInterface interface {
	GetAllTask(ctx context.Context) error
	CreateTask(ctx context.Context) error
}

type taskService struct {
	taskRepo repository.TaskRepositoryInterface
}

func NewTaskService(taskRepo repository.TaskRepositoryInterface) *taskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (s *taskService) GetAllTask(ctx context.Context) error {
	return nil
}

func (s *taskService) CreateTask(ctx context.Context) error {
	return nil
}

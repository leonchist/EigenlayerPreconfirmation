package service

import (
	"context"

	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/object"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/repository"
)

type TaskServiceInterface interface {
	GetAllTask() error
	CreateTask(ctx context.Context, req object.TaskObjRequest) error
}

type taskService struct {
	taskRepo repository.TaskRepositoryInterface
}

func NewTaskService(taskRepo repository.TaskRepositoryInterface) *taskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (s *taskService) GetAllTask() error {
	return nil
}

func (s *taskService) CreateTask(ctx context.Context, req object.TaskObjRequest) error {
	err := s.taskRepo.Create(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

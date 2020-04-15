package service

import (
	"context"

	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/object"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/repository"
)

type TaskServiceInterface interface {
	GetAllTask(ctx context.Context) (response []object.TaskObjResponse, err error)
	CreateTask(ctx context.Context, req object.TaskObjRequest) error
	UpdateTask(ctx context.Context, req object.TaskUpdateObjRequest) error
}

type taskService struct {
	taskRepo repository.TaskRepositoryInterface
}

func NewTaskService(taskRepo repository.TaskRepositoryInterface) *taskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (s *taskService) GetAllTask(ctx context.Context) (response []object.TaskObjResponse, err error) {
	resp, err := s.taskRepo.GetAll(ctx)
	if err != nil {
		return response, err
	}

	for _, v := range resp {
		obj := object.TaskObjResponse{
			ID:        int(v.ID),
			Name:      v.Name,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		response = append(response, obj)
	}
	return response, nil
}

func (s *taskService) CreateTask(ctx context.Context, req object.TaskObjRequest) error {
	err := s.taskRepo.Create(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s *taskService) UpdateTask(ctx context.Context, req object.TaskUpdateObjRequest) error {
	err := s.taskRepo.Update(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"context"

	"gitlab.warungpintar.co/back-end/libwp/database"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/model"
)

type TaskRepositoryInterface interface {
	GetAll(ctx context.Context) ([]model.Task, error)
	Create(ctx context.Context, req model.Task) error
}
type taskRepository struct {
	db *database.Store
}

func NewTaskRepository(db *database.Store) *taskRepository  {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) GetAll(ctx context.Context) ([]model.Task, error) {
	return []model.Task{}, nil
}

func (r *taskRepository) Create(ctx context.Context, req model.Task) error {
	return nil
}
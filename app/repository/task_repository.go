package repository

import (
	"context"
	"fmt"

	"gitlab.warungpintar.co/back-end/libwp/database"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/model"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/object"
)

type TaskRepositoryInterface interface {
	GetAll(ctx context.Context) ([]model.Task, error)
	Create(ctx context.Context, req object.TaskObjRequest) error
}
type taskRepository struct {
	db *database.Store
}

func NewTaskRepository(db *database.Store) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) GetAll(ctx context.Context) ([]model.Task, error) {
	return []model.Task{}, nil
}

func (r *taskRepository) Create(ctx context.Context, req object.TaskObjRequest) error {
	queryStr := fmt.Sprintf(`INSERT INTO tasks
	(name, status)
	VALUES
	(?, ?)
	`)
	_, err := r.db.Master.ExecContext(ctx, queryStr, req.Name,
		req.Status)

	if err != nil {
		return err
	}
	return nil

}

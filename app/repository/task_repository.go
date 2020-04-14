package repository

import (
	"context"
	"fmt"

	"gitlab.warungpintar.co/back-end/libwp/database"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/model"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/object"
)

type TaskRepositoryInterface interface {
	GetAll(ctx context.Context) (data []model.Task, err error)
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

func (r *taskRepository) GetAll(ctx context.Context) (data []model.Task, err error) {
	queryStr := fmt.Sprintf(`
	SELECT 
		id, name, status, created_at, updated_at
	FROM tasks`)

	rows, err := r.db.Slave.QueryContext(ctx, queryStr)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		model := model.Task{}
		if err = rows.Scan(&model.ID, &model.Name, &model.Status, &model.CreatedAt, &model.UpdatedAt); err != nil {
			return data, err
		}
		data = append(data, model)
	}
	return data, nil
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

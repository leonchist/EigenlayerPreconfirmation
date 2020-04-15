package service

import (
	"context"
	"testing"

	. "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/model"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/object"
	mockRepositories "gitlab.warungpintar.co/sharing-session/practice-go-unit-test/gen/mocks"
)

func Test_Get_Task(t *testing.T) {
	t.Run("Should get task data", func(t *testing.T) {
		ctrl := NewController(t)
		defer ctrl.Finish()

		mock := mockRepositories.NewMockTaskRepositoryInterface(ctrl)
		expected := []model.Task{
			model.Task{
				ID:        0,
				Name:      "Testing 0",
				Status:    "todo",
				CreatedAt: "2020-04-14T22:54:15Z",
				UpdatedAt: "2020-04-14T22:54:15Z",
			}, model.Task{
				ID:        1,
				Name:      "Testing 1",
				Status:    "Doing",
				CreatedAt: "2020-04-14T22:54:15Z",
				UpdatedAt: "2020-04-14T22:54:15Z",
			},
			model.Task{
				ID:        2,
				Name:      "Testing 0",
				Status:    "done",
				CreatedAt: "2020-04-14T22:54:15Z",
				UpdatedAt: "2020-04-14T22:54:15Z",
			},
		}

		mock.EXPECT().GetAll(context.Background()).Return(expected, nil).Times(1)

		service := taskService{
			taskRepo: mock,
		}
		result, err := service.GetAllTask(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, int(expected[0].ID), result[0].ID)
		assert.Equal(t, expected[0].Name, result[0].Name)
		assert.Equal(t, expected[0].Status, result[0].Status)
		assert.Equal(t, expected[0].CreatedAt, result[0].CreatedAt)
		assert.Equal(t, expected[0].UpdatedAt, result[0].UpdatedAt)
	})

}

func Test_Create_Task(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	mock := mockRepositories.NewMockTaskRepositoryInterface(ctrl)
	expected := object.TaskObjRequest{
		Name:   "okky",
		Status: "todo",
	}
	mock.EXPECT().Create(context.Background(), expected).Times(1)

	service := taskService{
		taskRepo: mock,
	}
	err := service.CreateTask(context.Background(), expected)

	assert.Nil(t, err)
}

func Test_Update_Task(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	mock := mockRepositories.NewMockTaskRepositoryInterface(ctrl)
	expected := object.TaskUpdateObjRequest{
		ID:     1,
		Name:   "okky",
		Status: "todo",
	}
	mock.EXPECT().Update(context.Background(), expected).Times(1)

	service := taskService{
		taskRepo: mock,
	}
	err := service.UpdateTask(context.Background(), expected)

	assert.Nil(t, err)
}

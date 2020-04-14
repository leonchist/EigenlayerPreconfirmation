package service

import (
	"context"
	"testing"

	. "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/object"
	mockRepositories "gitlab.warungpintar.co/sharing-session/practice-go-unit-test/gen/mocks"
)

func Test_Create_Task(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	// ctx := context.DeadlineExceeded
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

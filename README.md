# practice-go-unit-test using GOMOCK
Implement gomock

## Mocking
Merupakan cara pengujian perilaku sebuah service dengan cara membuat object palsu sehingga pengujian nya dapat terisolasi dari object yg sebenernya 

## apa itu GOMOCK ?
Gomock sendiri adalah sebuah framework untuk mocking di bahasa pemrograman Golang
- gomock - [gomock](https://github.com/golang/mock)

### Contoh Case
#### adapun kita memiliki sebuah service bernama `task_service.go` yg memiliki kebutuhan untuk berkomunikasi dengan `repository.TaskRepositoryInterface`
```
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

```
### Mocking Repository
Untuk mengisolasi sebuah repository `task_repository.go` yaitu dengan melakukan `mocking` terhadap `contract` atau `interface` yg ada
```
type TaskRepositoryInterface interface {
	GetAll(ctx context.Context) (data []model.Task, err error)
	Create(ctx context.Context, req object.TaskObjRequest) error
	Update(ctx context.Context, req object.TaskUpdateObjRequest) error
}
```

Dengan menggunakan `mockgen --source=task_repository.go`
Maka hasil generate mock nya adalah seperti ini 
```
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/model"
	object "gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/object"
)

// MockTaskRepositoryInterface is a mock of TaskRepositoryInterface interface
type MockTaskRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepositoryInterfaceMockRecorder
}

// MockTaskRepositoryInterfaceMockRecorder is the mock recorder for MockTaskRepositoryInterface
type MockTaskRepositoryInterfaceMockRecorder struct {
	mock *MockTaskRepositoryInterface
}

// NewMockTaskRepositoryInterface creates a new mock instance
func NewMockTaskRepositoryInterface(ctrl *gomock.Controller) *MockTaskRepositoryInterface {
	mock := &MockTaskRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockTaskRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskRepositoryInterface) EXPECT() *MockTaskRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockTaskRepositoryInterface) GetAll(ctx context.Context) ([]model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockTaskRepositoryInterfaceMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTaskRepositoryInterface)(nil).GetAll), ctx)
}

// Create mocks base method
func (m *MockTaskRepositoryInterface) Create(ctx context.Context, req object.TaskObjRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockTaskRepositoryInterfaceMockRecorder) Create(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskRepositoryInterface)(nil).Create), ctx, req)
}

// Update mocks base method
func (m *MockTaskRepositoryInterface) Update(ctx context.Context, req object.TaskUpdateObjRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockTaskRepositoryInterfaceMockRecorder) Update(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskRepositoryInterface)(nil).Update), ctx, req)
}
```

### Implementasi unit test untuk service nya `task_service_test.go`

```
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
```
Dengan seperti ini service yang membutuhkan object dari sebuah repository bisa di inject dengan mocking supaya terisolasi dari object yg sebenernya tetapi dengan perilaku yang sama

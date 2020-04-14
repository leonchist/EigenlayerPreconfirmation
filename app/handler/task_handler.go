package handler

import (
	"net/http"

	"gitlab.warungpintar.co/back-end/libwp/pkg/response"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/service"
)

type TaskHandlerInterface interface {
	GetTasks(r *http.Request) *response.JSONResponse
	CreateTask(r *http.Request) *response.JSONResponse
}

type taskHandler struct {
	taskSvc service.TaskServiceInterface
}

func NewTaskHandler(t service.TaskServiceInterface) *taskHandler {
	return &taskHandler{
		taskSvc: t,
	}
}

func (h *taskHandler) GetTasks(r *http.Request) *response.JSONResponse {
	resp := response.NewJSONResponse()
	return resp
}

func (h *taskHandler) CreateTask(r *http.Request) *response.JSONResponse {
	resp := response.NewJSONResponse()
	return resp
}

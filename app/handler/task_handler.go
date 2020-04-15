package handler

import (
	"encoding/json"
	"net/http"

	"gitlab.warungpintar.co/back-end/libwp/pkg/response"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/object"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/service"
)

type TaskHandlerInterface interface {
	GetTasks(r *http.Request) *response.JSONResponse
	CreateTask(r *http.Request) *response.JSONResponse
	UpdateTask(r *http.Request) *response.JSONResponse
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

	ctx := r.Context()
	data, err := h.taskSvc.GetAllTask(ctx)
	if err != nil {
		resp.SetError(err)
		return resp
	}
	resp.Data = data
	resp.Message = "Success get task data"
	resp.Code = "200"
	return resp
}

func (h *taskHandler) CreateTask(r *http.Request) *response.JSONResponse {
	resp := response.NewJSONResponse()
	var request object.TaskObjRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		resp.SetError(err)
		return resp
	}

	ctx := r.Context()
	err := h.taskSvc.CreateTask(ctx, request)
	if err != nil {
		resp.SetError(err)
		return resp
	}

	resp.Message = "Success add task"
	resp.Code = "201"
	return resp
}

func (h *taskHandler) UpdateTask(r *http.Request) *response.JSONResponse {
	resp := response.NewJSONResponse()
	var request object.TaskUpdateObjRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		resp.SetError(err)
		return resp
	}

	ctx := r.Context()
	err := h.taskSvc.UpdateTask(ctx, request)
	if err != nil {
		resp.SetError(err)
		return resp
	}

	resp.Message = "Success update task"
	resp.Code = "201"
	return resp
}

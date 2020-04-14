package router

import (
	"log"

	"gitlab.warungpintar.co/back-end/libwp/router"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/handler"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/config"

	grace "gitlab.warungpintar.co/back-end/libwp/httputil"
)

type Router struct {
	options       *router.MyRouter
	ListenAddress string
	listenErrCh   chan error
}

func NewRouter(
	cfg config.MainConfig,
	task handler.TaskHandlerInterface,
) *Router {
	r := router.New(&router.Options{
		Prefix:  cfg.Server.BasePath,
		Timeout: 10,
	})

	// Task
	r.GET("/tasks", task.GetTasks)

	return &Router{
		options: r,
	}

}

func (r *Router) Run() {
	log.Printf("API Listening on %s", r.ListenAddress)
	r.listenErrCh <- grace.Serve(r.ListenAddress, router.WrapperHandler(), 10, 10, 10)
}

func (r *Router) ListenError() <-chan error {
	return r.listenErrCh
}

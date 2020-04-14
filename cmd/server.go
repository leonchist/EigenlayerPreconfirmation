package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"gitlab.warungpintar.co/back-end/libwp/database"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/handler"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/repository"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/router"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/app/service"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/config"
)

func ServeCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "starting practice",
		Long:  `okky muhamad budiman practice unit testing`,
		Run: func(cmd *cobra.Command, args []string) {
			serve()
		},
	}
}

func serve() {
	cfg := &config.MainConfig{}
	config.ReadConfig(cfg, "main")

	// Database
	db := database.New(cfg.Database, database.DriverMySQL)
	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)

	taskHandler := handler.NewTaskHandler(taskService)

	server := router.NewRouter(*cfg, taskHandler)
	server.ListenAddress = cfg.Server.Port

	go server.Run()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Infoln("Exiting gracefully...")
	case err := <-server.ListenError():
		log.Errorln("Error starting web server, exiting gracefully:", err)
	}

}

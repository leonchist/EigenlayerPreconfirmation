package cmd

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
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
	// db := database.New(cfg.Database, database.DriverMySQL)

	log.Infof("Starting at port:%s", cfg.Server.Port)

	// if err := srv.Serve(lis); err != nil {
	// 	log.Errorln("Error starting material grpc server, exiting gracefully:", err)
	// }

	http.ListenAndServe(cfg.Server.Port, nil)

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Infoln("Exiting gracefully...")
	}

}

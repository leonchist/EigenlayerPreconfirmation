package main

import (
	"github.com/spf13/cobra"
	"gitlab.warungpintar.co/sharing-session/practice-go-unit-test/cmd"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "learn-testing",
	}

	rootCmd.AddCommand(cmd.ServeCMD())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

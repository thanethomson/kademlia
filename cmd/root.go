package cmd

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kademlia",
	Short: "Runs a Kademlia node for experimentation purposes.",
	Run: func(cmd *cobra.Command, args []string) {
		logging.SetBackend(
			logging.NewBackendFormatter(
				logging.NewLogBackend(os.Stdout, "", 0),
				logging.MustStringFormatter(
					`%{color}%{time:2006-01-02T15:04:05Z07:00} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
				),
			),
		)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

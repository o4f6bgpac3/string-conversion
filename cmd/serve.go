package cmd

import (
	"github.com/o4f6bgpac3/string-conversion/internal/config"
	"github.com/o4f6bgpac3/string-conversion/internal/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		s := server.New(cfg)
		return s.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

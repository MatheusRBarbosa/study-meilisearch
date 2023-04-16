package cmd

import (
	"github.com/spf13/cobra"
	"squad10x.com.br/boilerplate/api"
)

var (
	apiCommand = &cobra.Command{
		Use:   "api",
		Short: "Initializes HTTP server",
		Long:  "Initializes the codebase to serve HTTP API server",
		RunE:  apiExecute,
	}
)

func init() {
	rootCmd.AddCommand(apiCommand)
}

func apiExecute(_ *cobra.Command, _ []string) error {
	api.StartHttpServer()
	return nil
}

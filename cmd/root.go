package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "10x",
	Short: "10x base command",
	Long:  "10x base command",
}

func Execute() {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal("Unexpected error:", err)
		}
	}()

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("Error on Execute:", err)
	}
}

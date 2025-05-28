/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/booscaaa/clean-go/pkg-by-feature/pkg/rest"
	"github.com/spf13/cobra"

	"github.com/booscaaa/initializers/postgres"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		database := postgres.Initialize(
			cmd.Context(),
			"postgres://postgres:postgres@db:5432/example?sslmode=disable",
		) //

		rest.Initialize(database)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

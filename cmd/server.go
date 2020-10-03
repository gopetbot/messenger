package cmd

import (
	"github.com/gopetbot/messenger/pet"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "server",
	Short: "short description",
	RunE: func(cmd *cobra.Command, args []string) error {
		app := pet.Run()

		if err := app.Server(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

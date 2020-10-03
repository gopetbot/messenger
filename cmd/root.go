package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	configFileFlag string
)

var rootCmd = &cobra.Command{
	Use:   "pet",
	Short: "A brief description of your application",
	Long:  "A longer description: Web Application for track news related by covid",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
//TODO implements init config

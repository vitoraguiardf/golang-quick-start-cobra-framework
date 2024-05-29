package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Cristina",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cristina v0.1")
	},
}

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cristina",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("My first cli app in golang!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

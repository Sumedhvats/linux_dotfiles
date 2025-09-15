package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a command to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		strings.Join(args," ")
		println()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

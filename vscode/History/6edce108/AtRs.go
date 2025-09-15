package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete a particular command",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := []string{}
		for _, arg := range args {
			tasks = append(tasks, strconv(arg))
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}

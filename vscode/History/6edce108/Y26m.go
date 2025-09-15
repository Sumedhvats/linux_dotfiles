package cmd

import (
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete a particular command",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := []string{}
		for i, arg := range args {
			tasks = append(tasks, )
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}

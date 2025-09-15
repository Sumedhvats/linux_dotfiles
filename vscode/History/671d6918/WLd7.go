package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a command to your task list",
	Run: func(cmd *cobra.Command, args []string) {
	for i,arg:= range args{

	}
	},
}
func init(){
	rootCmd.AddCommand(addCmd)
}
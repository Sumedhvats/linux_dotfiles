
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "r brief description of your command",
Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}

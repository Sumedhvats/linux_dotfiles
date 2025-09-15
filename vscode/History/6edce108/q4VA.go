package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete a particular command",
Run: func(cmd *cobra.Command, args []string) {
tasks :={}
	for i,arg:=range args{

}
},
}

func init() {
	rootCmd.AddCommand(doCmd)
}

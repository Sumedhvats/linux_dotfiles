package cmd

import (
	"strconv"
	"text/template/parse"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete a particular command",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := []int{}
		for _, arg := range args {
			curr,err:=strconv.Atoi(arg)
			if err!=nil{
				panic(err)
			}
			tasks=append(tasks, curr)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}

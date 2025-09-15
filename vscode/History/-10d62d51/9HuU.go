/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/Sumedhvats/task/cmd"
)

func main() {
	home,_:= os.UserHomeDir()
	cmd.Execute()
}

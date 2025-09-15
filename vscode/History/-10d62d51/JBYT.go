/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"
	"path/filepath"

	"github.com/Sumedhvats/task/cmd"
)

func main() {
	home,_:= os.UserHomeDir()
	filepath.Join(home,"tasks.db")
	cmd.Execute()
}

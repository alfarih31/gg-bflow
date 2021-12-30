package main

import (
	"os"
)

var (
	cfgFile   string
	targetDir string
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

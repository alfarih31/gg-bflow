package main

import (
	"fmt"
	app "github.com/alfarih31/gg-bflow"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Version of GG-BFlow",
		Long:  "Show Version of GG-BFlow",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("GG-BFlow Version: %s\n", app.Version)
		},
	}

	rootCmd.AddCommand(cmd)
}

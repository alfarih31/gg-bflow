package main

import (
	gg_bflow "github.com/alfarih31/gg-bflow/pkg/gg-bflow"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/logger"
	"github.com/spf13/cobra"
)

func serve() (err error) {
	err = gg_bflow.Init(cfgFile)
	if err != nil {
		return err
	}
	gg_bflow.Start()

	return nil
}

func init() {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve GG BFlow",
		PreRun: func(cmd *cobra.Command, args []string) {
			logger.Log.Info("Serving...")
		},
		RunE: func(c *cobra.Command, args []string) (err error) {
			err = serve()
			if err != nil {
				return
			}

			return
		},
	}

	f := cmd.Flags()
	f.StringVarP(&cfgFile, "env", "e", "", `.env file for configuration (default "System Wide Environment")`)

	rootCmd.AddCommand(cmd)
}

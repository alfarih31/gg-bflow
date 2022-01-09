package main

import (
	"fmt"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/logger"
	"github.com/spf13/cobra"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

type files []string

func (f files) String() string {
	o := "Files:\n"
	for i, s := range f {
		o += fmt.Sprintf("  %d. %s\n", i+1, filepath.Base(s))
	}

	return o
}

func getSourceDir() (runDir string, err error) {
	_, ex, _, ok := runtime.Caller(0)
	if !ok {
		err = fmt.Errorf("%s", "No runtime!")
		return
	}

	runDir = filepath.Dir(ex)

	runDir = filepath.Join(filepath.Dir(ex), "../../")
	return
}

func listFiles(s string) (files, error) {
	var out []string
	err := filepath.Walk(s, func(path string, info fs.FileInfo, err error) error {
		if path == s {
			return nil
		}

		out = append(out, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return out, nil
}

func copyFile(s string, d string) (err error) {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(d, data, 0755)
	if err == nil {
		logger.Log.Info(fmt.Sprintf(`Export %s completed`, filepath.Base(s)), map[string]string{"targetPath": d})
	}
	return
}

func copyFiles(ss []string, destDir string) (err error) {
	for _, s := range ss {
		t := filepath.Join(destDir, filepath.Base(s))
		err = copyFile(s, t)
		if err != nil {
			return
		}
	}

	return
}

func exportEnvCmd(rc *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "env",
		Short: "Export .env.example",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			runDir, err := getSourceDir()
			if err != nil {
				return
			}

			envFile := filepath.Join(runDir, ".env.example")
			logger.Log.Info(fmt.Sprintf("Copying .env.example from: %s ...", envFile))
			err = copyFile(envFile, filepath.Join(targetDir, ".env.example"))
			return
		},
	}

	rc.AddCommand(cmd)
}

func exportProtoCmd(rc *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "proto",
		Short: "Export gRPC Proto",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			runDir, err := getSourceDir()
			if err != nil {
				return
			}

			pbDir := filepath.Join(runDir, "proto")
			logger.Log.Info(fmt.Sprintf("Listing gRPC Proto Files on: %s ...", pbDir))
			ps, err := listFiles(pbDir)
			fmt.Println(ps)

			err = copyFiles(ps, targetDir)
			return
		},
	}

	rc.AddCommand(cmd)
}

func init() {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export GG-BBFlow API files, such: .env.example & gRPC proto",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger.Log.Infof("Run %s...", cmd.CommandPath())
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			logger.Log.Infof("%s completed", cmd.CommandPath())
		},
	}

	f := cmd.PersistentFlags()
	f.StringVarP(&targetDir, "out-dir", "o", "./", "Target export directory")

	exportEnvCmd(cmd)
	exportProtoCmd(cmd)

	rootCmd.AddCommand(cmd)
}

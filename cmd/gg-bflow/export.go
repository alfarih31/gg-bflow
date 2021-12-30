package main

import (
	"github.com/spf13/cobra"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getRuntimeDir() (runDir string, err error) {
	ex, err := os.Executable()
	if err != nil {
		return
	}

	runDir = filepath.Dir(ex)
	return
}

func listFiles(s string) (files []string, err error) {
	err = filepath.Walk(s, func(path string, info fs.FileInfo, err error) error {
		if path == s {
			return nil
		}

		files = append(files, path)
		return nil
	})

	return
}

func copyFile(s string, d string) (err error) {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(d, data, 0755)
	return
}

func copyFiles(ss []string, destDir string) (err error) {
	for _, s := range ss {
		err = copyFile(s, filepath.Join(destDir, filepath.Base(s)))
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
			runDir, err := getRuntimeDir()
			if err != nil {
				return
			}

			err = copyFile(filepath.Join(runDir, ".env.example"), filepath.Join(targetDir, ".env.example"))
			return
		},
	}

	rc.AddCommand(cmd)
}

func exportPBCmd(rc *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "pb",
		Short: "Export gRPC Generated PB",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			runDir, err := getRuntimeDir()
			if err != nil {
				return
			}

			files, err := listFiles(filepath.Join(runDir, "api/grpc"))

			err = copyFiles(files, targetDir)
			return
		},
	}

	rc.AddCommand(cmd)
}

func init() {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export GG-BBFlow API files, such: .env.example & gRPC Generated PB",
	}

	f := cmd.PersistentFlags()
	f.StringVarP(&targetDir, "out-dir", "o", "./", "Target export directory")

	exportEnvCmd(cmd)
	exportPBCmd(cmd)

	rootCmd.AddCommand(cmd)
}

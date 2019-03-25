package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash [flags] DIR",
	Short: "Creates a list of file hashes",
	Run:   runHash,
}

func runHash(cmd *cobra.Command, args []string) {
	path := getPath(args)
	err := checkPath(path)
	if err != nil {
		printFatal(err.Error())
	}
	storage, _ := cmd.Flags().GetString("storage")
	hash(path, filepath.ToSlash(storage))
}

func hash(path string, storage string) {
	storageName := storage
	if storageName == "" {
		storageName = "(not set)"
	}
	printStr(fmt.Sprintf("Start hashing. Directory: %s / Storage: %s", path, storageName))
	dir := HashDir(path, hashFunc)
	s := NewCSVStorage(dir, storage)
	err := exportDir(s)
	if err != nil {
		printFatal(fmt.Sprintf("An error occurred during hashing. Error: %v", err.Error()))
	}
	printStr(fmt.Sprintf("Hashing completed. Result saved: %s", s.storage))
}

func init() {
	rootCmd.AddCommand(hashCmd)
	hashCmd.Flags().StringP("storage", "s", "", "Storage path")
}

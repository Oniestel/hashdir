package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check [flags] DIR",
	Short: "Checks directory for changes and displays them",
	Run:   runCheck,
}

func runCheck(cmd *cobra.Command, args []string) {
	path := getPath(args)
	err := checkPath(path)
	if err != nil {
		printFatal(err.Error())
	}
	storage, _ := cmd.Flags().GetString("storage")
	err = checkStorage(storage)
	if err != nil {
		printFatal(err.Error())
	}
	check(path, filepath.ToSlash(storage))
}

func check(path string, storage string) {
	printStr(fmt.Sprintf("Check started. Directory: %s / Storage: %s", path, storage))
	dir := HashDir(path, hashFunc)
	oldDir := NewDir(path)
	err := importDir(NewCSVStorage(oldDir, storage))
	if err != nil {
		printFatal(fmt.Sprintf("An error occurred during check. Error: %v", err.Error()))
	}
	added, modified, deleted := CompareDirs(oldDir, dir)
	printStr("Check completed.")
	printAddedFiles(added)
	printModifiedFiles(modified)
	printDeletedFiles(deleted)
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringP("storage", "s", "", "Storage path")
}

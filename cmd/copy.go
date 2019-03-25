package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy [flags] DIR",
	Short: "Copies changes to dest",
	Run:   runCopy,
}

func runCopy(cmd *cobra.Command, args []string) {
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
	dest, _ := cmd.Flags().GetString("dest")
	err = checkDest(dest)
	if err != nil {
		printFatal(err.Error())
	}
	copy(path, filepath.ToSlash(storage), filepath.ToSlash(dest))
}

func copy(path string, storage string, dest string) {
	printStr(fmt.Sprintf("Copy started. Directory: %s / Storage: %s / Destination: %s", path, storage, dest))
	dir := HashDir(path, hashFunc)
	oldDir := NewDir(path)
	err := importDir(NewCSVStorage(oldDir, storage))
	if err != nil {
		printFatal(fmt.Sprintf("An error occurred during copy. Error: %v", err.Error()))
	}
	added, modified, _ := CompareDirs(oldDir, dir)
	copyFiles(dest, added)
	copyFiles(dest, modified)
	printStr("Copy completed.")
}

func copyFiles(dest string, files map[string]*File) {
	for _, file := range files {
		err := CopyFile(file, dest)
		if err != nil {
			printStr("Couldn't copy file: " + file.Path)
		}
		printStr(file.Path)
	}
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().StringP("storage", "s", "", "Storage path")
	copyCmd.Flags().StringP("dest", "d", "", "Dest path")
}

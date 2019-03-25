package cmd

import (
	"fmt"
	"os"
)

func printStr(str string) {
	fmt.Println(str)
}

func printErr(err error) {
	fmt.Println(err)
}

func printFatal(str string) {
	fmt.Println(str)
	os.Exit(1)
}

func printAddedFiles(files map[string]*File) {
	printStr("Added files:")
	printFiles(files)
}

func printDeletedFiles(files map[string]*File) {
	printStr("Deleted files:")
	printFiles(files)
}

func printModifiedFiles(files map[string]*File) {
	printStr("Modified files:")
	printFiles(files)
}

func printFiles(files map[string]*File) {
	for _, file := range files {
		printStr(file.Path)
	}
}

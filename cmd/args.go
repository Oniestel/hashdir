package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

//ValidateError arguments validation error
type ValidateError struct {
	Msg string
}

func (r ValidateError) Error() string {
	return fmt.Sprintf("%s", r.Msg)
}

func getPath(args []string) (path string)  {
	if len(args) > 0 {
		path = filepath.ToSlash(args[0])
	}
	return
}

func checkPath(path string) error {
	if path == "" {
		return &ValidateError{"Path is not specified."}
	}
	if !dirExist(path) {
		return &ValidateError{"Target directory doesn't exist."}
	}
	return nil
}

func checkStorage(storage string) error {
	if storage == "" {
		return &ValidateError{"Storage path can't be empty."}
	}
	return nil
}

func checkDest(dest string) error {
	if dest == "" {
		return &ValidateError{"Destination path can't be empty."}
	}
	if !dirExist(dest) {
		return &ValidateError{"Destination directory doesn't exist."}
	}
	return nil
}

func dirExist(path string) bool {
	info, err := os.Stat(filepath.FromSlash(path))
	if err != nil {
		return false
	}
	if !info.IsDir() {
		return false
	}
	return true
}

package cmd

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//File describes contained in the dir
type File struct {
	Path string
	Hash string
}

//GetPathHash return hash of the path
func (r *File) GetPathHash(dirPath string) string {
	path := strings.Replace(r.Path, dirPath, "", 1)
	return getSHA1(path)
}

//FillHash fills the hash in the file structure
func (r *File) FillHash(hashFunc HashFunc) error {
	var err error
	r.Hash, err = hashFunc(r.Path)
	return err
}

//NewFile create new File
func NewFile(path string, hash string) *File {
	return &File{path, hash}
}

//SameFiles compares file hashes
func SameFiles(old *File, new *File) bool {
	return old.Hash == new.Hash
}

//CopyFile copies existing file to dest
func CopyFile(file *File, dest string) error {
	srcFile, err := os.Open(filepath.FromSlash(file.Path))
	if err != nil {
		return err
	}
	defer srcFile.Close()
	destFilePath := replaceParentDir(file.Path, "/", dest)
	os.MkdirAll(filepath.Dir(destFilePath), 0644)
	destFile, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, srcFile)
	return err
}

func replaceParentDir(path string, pathSep string, newDir string) string {
	parts := strings.Split(path, pathSep)
	if len(parts) == 0 {
		return newDir
	}
	return strings.Replace(path, parts[0], newDir, 1)
}

func getSHA1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

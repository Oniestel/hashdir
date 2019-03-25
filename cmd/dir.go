package cmd

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

//HashFunc describes hash function type
type HashFunc func(path string) (string, error)

//Dir contains directory info and files
type Dir struct {
	Path  string
	Files map[string]*File
	queue chan *File
	wg    sync.WaitGroup
}

//Init new dir
func (r *Dir) Init(path string) {
	r.Path = path
	r.Files = make(map[string]*File)
}

//FillFiles reads files of a directory
func (r *Dir) FillFiles() {
	err := filepath.Walk(r.Path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		slashPath := filepath.ToSlash(path)
		file := NewFile(slashPath, "")
		r.Files[file.GetPathHash(r.Path)] = file
		return nil
	})
	if err != nil {
		printErr(err)
	}
}

//HashFiles creates file hashes
func (r *Dir) HashFiles(hashFunc HashFunc) {
	r.queue = make(chan *File)
	r.runWorkers(hashFunc)
	r.processFiles()
	r.wg.Wait()
	close(r.queue)
}

func (r *Dir) runWorkers(hashFunc HashFunc) {
	for i := 0; i < runtime.NumCPU(); i++ {
		go r.hashWorker(hashFunc)
	}
}

func (r *Dir) processFiles() {
	for _, file := range r.Files {
		r.wg.Add(1)
		r.queue <- file
	}
}

func (r *Dir) hashWorker(hashFunc HashFunc) {
	for {
		file, more := <-r.queue
		if !more {
			return
		}
		err := file.FillHash(hashFunc)
		r.wg.Done()
		if err != nil {
			printErr(err)
			continue
		}
	}
}

//NewDir creates new Dir
func NewDir(path string) *Dir {
	dir := &Dir{}
	dir.Init(path)
	return dir
}

//HashDir hash files using hash function
func HashDir(path string, hashFunc HashFunc) *Dir {
	dir := NewDir(path)
	dir.FillFiles()
	dir.HashFiles(hashFunc)
	return dir
}

//CompareDirs compare two dirs. Return added, modified and deleted files.
func CompareDirs(old *Dir, new *Dir) (added map[string]*File, modified map[string]*File, deleted map[string]*File) {
	oldFiles := copyFilesMap(old.Files)
	newFiles := copyFilesMap(new.Files)
	modified = make(map[string]*File)
	for key, oldFile := range oldFiles {
		if newFile, ok := newFiles[key]; ok {
			delete(oldFiles, key)
			delete(newFiles, key)
			if !SameFiles(oldFile, newFile) {
				modified[key] = oldFile
			}
		}
	}
	deleted = oldFiles
	added = newFiles
	return
}

func copyFilesMap(src map[string]*File) map[string]*File {
	dest := make(map[string]*File)
	for key, value := range src {
		dest[key] = value
	}
	return dest
}

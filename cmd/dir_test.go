package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNewDir(t *testing.T) {
	dir := NewDir("test")
	if dir.Path != "test" {
		t.Fail()
	}
}

func TestDir_Init(t *testing.T) {
	TestNewDir(t)
}

func TestDir_FillFiles(t *testing.T) {
	defer removeTestDir()
	path := createTestDir(t)
	dir := NewDir(path)
	dir.FillFiles()
	if len(dir.Files) == 0 {
		t.Fail()
	}
}

func TestDir_HashFiles(t *testing.T) {
	defer removeTestDir()
	path := createTestDir(t)
	dir := NewDir(path)
	dir.FillFiles()
	dir.HashFiles(testHashFunc)
	for _, file := range dir.Files {
		if file.Hash != "hash" {
			t.Fail()
		}
	}
}

func TestHashDir(t *testing.T) {
	defer removeTestDir()
	path := createTestDir(t)
	dir := HashDir(path, testHashFunc)
	for _, file := range dir.Files {
		if file.Hash != "hash" {
			t.Fail()
		}
	}
}

func TestCompareDirs(t *testing.T) {
	defer removeTestDir()
	dirPath := createTestDir(t)
	dir := HashDir(dirPath, hashFunc)
	removeTestDir()
	dirPath = createTestDirModified(t)
	modifiedDir := HashDir(dirPath, hashFunc)
	a, m, d := CompareDirs(dir, modifiedDir)
	if len(a) == 0 {
		t.Error("Expected non-zero len for added files")
	}
	if len(m) == 0 {
		t.Error("Expected non-zero len for modified files")
	}
	if len(d) == 0 {
		t.Error("Expected non-zero len for deleted files")
	}
}

func testHashFunc(path string) (s string, e error) {
	return "hash", nil
}

func removeTestDir() {
	os.RemoveAll("test")
}

func createTestDir(t *testing.T) string {
	path := "test"
	err := os.Mkdir(path, 0644)
	if err != nil {
		t.Fail()
	}
	err = ioutil.WriteFile(path+"/test", []byte(""), 0644)
	if err != nil {
		t.Fail()
	}
	err = ioutil.WriteFile(path+"/test2", []byte(""), 0644)
	if err != nil {
		t.Fail()
	}
	return path
}

func createTestDirModified(t *testing.T) string {
	path := "test"
	err := os.Mkdir(path, 0644)
	if err != nil {
		t.Fail()
	}
	err = ioutil.WriteFile(path+"/test", []byte("content"), 0644)
	if err != nil {
		t.Fail()
	}
	err = ioutil.WriteFile(path+"/test3", []byte(""), 0644)
	if err != nil {
		t.Fail()
	}
	return path
}

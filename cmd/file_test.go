package cmd

import (
	"testing"
)

func TestNewFile(t *testing.T) {
	file := NewFile("path", "hash")
	if file.Path != "path" {
		t.Fail()
	}
	if file.Hash != "hash" {
		t.Fail()
	}
}

func TestFile_GetPathHash(t *testing.T) {
	file := NewFile("test", "")
	hash := file.GetPathHash("test")
	if hash == "" {
		t.Fail()
	}
}

func TestFile_FillHash(t *testing.T) {
	file := NewFile("test", "")
	err := file.FillHash(testHashFunc)
	if err != nil {
		t.Fail()
	}
	if file.Hash != "hash" {
		t.Fail()
	}
}

func TestSameFiles(t *testing.T) {
	file := NewFile("test", "hash")
	file2 := NewFile("test", "hash2")
	if SameFiles(file, file2) {
		t.Fail()
	}
	file2 = NewFile("test", "hash")
	if !SameFiles(file, file2) {
		t.Fail()
	}
}

func TestCopyFile(t *testing.T) {
	defer removeTestDir()
	createTestDir(t)
	file := NewFile("test/test", "")
	err := CopyFile(file, "test/test_copy")
	if err != nil {
		t.Fail()
	}
}

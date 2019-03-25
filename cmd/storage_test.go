package cmd

import (
	"os"
	"testing"
)

type TestStorage struct {
	dir *Dir
}

func (TestStorage) Import() error {
	return nil
}

func (TestStorage) Export() error {
	return nil
}

func NewTestStorage(dir *Dir) *TestStorage {
	return &TestStorage{dir: dir}
}

func TestExportDir(t *testing.T) {
	err := exportDir(NewTestStorage(NewDir("test")))
	if err != nil {
		t.Error("Expected nil for error")
	}
}

func TestImportDir(t *testing.T) {
	err := importDir(NewTestStorage(NewDir("test")))
	if err != nil {
		t.Error("Expected nil for error")
	}
}

func TestCSVStorage_Export(t *testing.T) {
	defer func() {
		os.Remove("test.csv")
		removeTestDir()
	}()
	createTestDir(t)
	d := NewDir("test")
	s := NewCSVStorage(d, "test.csv")
	err := s.Export()
	if err != nil {
		t.Error("Expected nil for error")
	}
	if _, err := os.Stat("test.csv"); os.IsNotExist(err) {
		t.Error("Exported file doesn't exist")
	}
}

func TestCSVStorage_Import(t *testing.T) {
	defer func() {
		os.Remove("test.csv")
		removeTestDir()
	}()
	createTestDir(t)
	d := HashDir("test", testHashFunc)
	s := NewCSVStorage(d, "test.csv")
	err := exportDir(s)
	if err != nil {
		t.Error("Expected nil for error")
	}
	d = NewDir("test")
	s = NewCSVStorage(d, "test.csv")
	err = s.Import()
	if err != nil {
		t.Error("Expected nil for error")
	}
	if len(d.Files) == 0 {
		t.Error("Expected non-zero len for files")
	}
}

func TestNewCSVStorage(t *testing.T) {
	d := NewDir("test")
	s := NewCSVStorage(d, "")
	if s.storage != "test.csv" {
		t.Error("Incorrect storage value")
	}
}

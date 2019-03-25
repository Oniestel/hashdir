package cmd

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"path"
	"path/filepath"
)

//StorageExporter exporter interface
type StorageExporter interface {
	Export() error
}

//StorageImporter importer interface
type StorageImporter interface {
	Import() error
}

func exportDir(exporter StorageExporter) error {
	return exporter.Export()
}

func importDir(importer StorageImporter) error {
	return importer.Import()
}

//CSVStorage is storage based on the csv files
type CSVStorage struct {
	dir     *Dir
	storage string
}

//Export dir
func (r *CSVStorage) Export() error {
	f, err := os.Create(r.storage)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	for _, file := range r.dir.Files {
		if file.Hash == "" {
			continue
		}
		w.Write([]string{file.Path, file.Hash})
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

//Import dir
func (r *CSVStorage) Import() error {
	f, err := os.Open(filepath.FromSlash(r.storage))
	if err != nil {
		return err
	}
	reader := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if len(record) < 2 {
			continue
		}
		file := NewFile(record[0], record[1])
		r.dir.Files[file.GetPathHash(r.dir.Path)] = file
	}
	return nil
}

//NewCSVStorage creates new CSVStorage
func NewCSVStorage(dir *Dir, storage string) *CSVStorage {
	if storage == "" {
		storage = path.Base(dir.Path) + ".csv"
	}
	return &CSVStorage{dir: dir, storage: storage}
}

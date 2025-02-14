package httpserver

import (
	"net/http"
	"path/filepath"
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	file, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	fileStatistics, err := file.Stat()
	if err != nil {
		if closeErr := file.Close(); closeErr != nil {
			return nil, closeErr
		}
		return nil, err
	}

	if fileStatistics.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			if closeErr := file.Close(); closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}

	return file, nil
}

package filesystem

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Filesystem struct {
}

type File struct {
	IsDir bool
	Name  string
	Path  string
}

func (e Filesystem) GetFileList(root string) []File {
	fileList := []File{}
	first := true
	maxDepth := strings.Count(root, "/") + 1
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// handle possible path err, just in case...
			return err
		}
		if strings.Count(path, string(os.PathSeparator)) > maxDepth {
			fmt.Println("skip", path)
			return fs.SkipDir
		}

		if first {
			first = false
			return nil
		}

		fileList = append(fileList, File{d.IsDir(), d.Name(), path})
		return nil
	})

	sort.Slice(fileList, func(i, j int) bool {
		return fileList[i].IsDir
	})

	return fileList
}

func (e Filesystem) Delete(path string) {
	currentWorkingDirectory, _ := os.Getwd()
	minDepth := strings.Count(currentWorkingDirectory, "/")
	pathDepth := strings.Count(path, "/")

	if pathDepth > minDepth {
		os.RemoveAll(path)
	}
}

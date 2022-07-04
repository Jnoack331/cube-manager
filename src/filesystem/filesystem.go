package filesystem

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

type Filesystem struct {
}

type File struct {
	IsDir    bool
	Name     string
	Path     string
	Selected bool
}

func (e Filesystem) GetFileList(root string) []File {
	if e.isInRootPath(root) == false {
		return []File{}
	}

	fileList := []File{}
	first := true
	maxDepth := strings.Count(root, "/") + 1
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// handle possible path err, just in case...
			return err
		}
		if strings.Count(path, string(os.PathSeparator)) > maxDepth {
			return fs.SkipDir
		}

		if first {
			first = false
			return nil
		}

		fileList = append(fileList, File{d.IsDir(), d.Name(), path, false})
		return nil
	})

	sort.Slice(fileList, func(i, j int) bool {
		return fileList[i].IsDir
	})

	return fileList
}

func (e Filesystem) Delete(path string) {
	if e.isInRootPath(path) {
		os.RemoveAll(path)
	}
}

func (e Filesystem) isInRootPath(filePath string) bool {
	currentWorkingDirectory, _ := os.Getwd()
	filePath = path.Clean(filePath)
	minDepth := strings.Count(currentWorkingDirectory, "/")
	pathDepth := strings.Count(filePath, "/")
	if pathDepth < minDepth {
		return false
	}

	return strings.HasPrefix(filePath, filePath)
}

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
	Info os.DirEntry
	Path string
	Back bool
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
			splitted := strings.Split(path, "/")
			splitted = splitted[:len(splitted)-1]
			newPath := strings.Join(splitted, "/")
			fileList = append(fileList, File{d, newPath, true})
			first = false
		} else {
			fileList = append(fileList, File{d, path, false})
		}

		return nil
	})

	first = true
	sort.Slice(fileList, func(i, j int) bool {
		if fileList[j].Back {
			return false
		}

		if fileList[i].Back {
			return true
		}

		return fileList[i].Info.IsDir()
	})

	return fileList
}

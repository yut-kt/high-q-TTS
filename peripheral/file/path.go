package file

import (
	"io/ioutil"
	"path/filepath"
)

func FetchPaths(dir string) ([]string) {
	return dirWalk(dir)
}

func dirWalk(dir string) ([]string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			walked := dirWalk(filepath.Join(dir, file.Name()))
			paths = append(paths, walked...)
		} else {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths
}

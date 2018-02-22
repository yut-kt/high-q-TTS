package file

import (
	"io/ioutil"
	"path/filepath"
)

func FetchPaths(dir string) ([]string, error) {
	return dirWalk(dir)
}

func dirWalk(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			walked, err := dirWalk(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			paths = append(paths, walked...)
		} else {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths, nil
}

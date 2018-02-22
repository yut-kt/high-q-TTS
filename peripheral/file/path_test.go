package file

import (
	"path/filepath"
	"testing"
)

func TestFetchFilePaths(t *testing.T) {
	filePaths, err := FetchFilePaths()
	if err != nil {
		t.Fatal(err)
	}
	if len(filePaths) != 2 {
		t.Errorf("The number of elements of filePaths is incorrect")
	}

	if filepath.Base(filePaths[0]) != "test_dir_file.txt" {
		t.Errorf("Files under the test directory can not be acquired")
	}

	if filepath.Base(filePaths[1]) != "test_file.txt" {
		t.Errorf("Files can not be acquired")
	}
}

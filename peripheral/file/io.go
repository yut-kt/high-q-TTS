package file

import (
	"os"
	"bufio"
)

func Scanner(fp *os.File) func() (bool, string, error) {
	scanner := bufio.NewScanner(fp)
	return func() (bool, string, error) {
		scannabled := scanner.Scan()
		if err := scanner.Err(); err != nil {
			return scannabled, "", err
		}
		return scannabled, scanner.Text(), nil
	}
}

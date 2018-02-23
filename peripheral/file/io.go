package file

import (
	"os"
	"bufio"
)

func Open(path string) *os.File {
	finp, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return finp
}

func CreateOpen(path string) *os.File {
	foutp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return foutp
}

func Close(fp *os.File) {
	if err := fp.Close(); err != nil {
		panic(err)
	}
}

func Scanner(fp *os.File) func() (bool, string) {
	scanner := bufio.NewScanner(fp)
	return func() (bool, string) {
		scannabled := scanner.Scan()
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		return scannabled, scanner.Text()
	}
}

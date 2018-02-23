package main

import (
	"github.com/yut-kt/high-q-TTS/env"
	"github.com/yut-kt/high-q-TTS/peripheral/file"
	"github.com/yut-kt/high-q-TTS/peripheral/analysis"
	"path/filepath"
)

func main() {
	finp := file.Open(env.AllUniqueBodyPath)
	defer file.Close(finp)
	foutp := file.CreateOpen(env.AllPhraseDivisionBodyPath)
	defer file.Close(foutp)

	scan := file.Scanner(finp)
	cabocha := analysis.NewCabocha()
	for scannabled, text := scan(); scannabled; scannabled, text = scan() {
		foutp.WriteString(cabocha.ParseToWakati(text) + "\n")
	}

	filePaths := file.FetchPaths(env.ClusterNormalDir)
	for _, filePath := range filePaths {
		finp := file.Open(filePath)
		scan := file.Scanner(finp)
		foutp := file.CreateOpen(env.ClusterPhraseDivisionDir + "/" + filepath.Base(filePath))

		for scannabled, text := scan(); scannabled; scannabled, text = scan() {
			foutp.WriteString(cabocha.ParseToWakati(text) + "\n")
		}

		file.Close(finp)
		file.Close(foutp)
	}
}

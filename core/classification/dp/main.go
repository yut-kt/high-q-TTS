package main

import (
	"github.com/yut-kt/high-q-TTS/env"
	"github.com/yut-kt/high-q-TTS/peripheral/analysis"
	"github.com/yut-kt/high-q-TTS/peripheral/file"
	"os"
	"path/filepath"
	"regexp"
)

type cluster map[string][]string

func main() {
	filePaths, err := file.FetchPaths(env.ClusterBaseDir)
	if err != nil {
		panic(err)
	}

	rep := regexp.MustCompile(`.txt$`)
	base := make(cluster, len(filePaths))
	for _, filePath := range filePaths {
		clusterName := filepath.Base(rep.ReplaceAllString(filePath, ""))

		finp, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		scan := file.Scanner(finp)
		for scannabled, text, err := scan(); scannabled; scannabled, text, err = scan() {
			if err != nil {
				panic(err)
			}
			base[clusterName] = append(base[clusterName], text)
		}

		if err := finp.Close(); err != nil {
			panic(err)
		}
	}
	println(len(base))

	cabocha := analysis.NewCabocha()
	println(cabocha.ParseToWakati("あなたとJava"))

}

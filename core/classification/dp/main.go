package main

import (
	"github.com/yut-kt/high-q-TTS/env"
	"github.com/yut-kt/high-q-TTS/peripheral/file"
	"path/filepath"
	"regexp"
	"github.com/yut-kt/high-q-TTS/core/classification/dp/levenshtein"
)

type cluster map[string][]string

func main() {
	filePaths := file.FetchPaths(env.ClusterPhraseDivisionDir)

	rep := regexp.MustCompile(`.txt$`)
	base := make(cluster, len(filePaths))
	for _, filePath := range filePaths {
		clusterName := filepath.Base(rep.ReplaceAllString(filePath, ""))

		finp := file.Open(filePath)

		scan := file.Scanner(finp)
		for scannabled, text := scan(); scannabled; scannabled, text = scan() {
			base[clusterName] = append(base[clusterName], text)
		}

		file.Close(finp)
	}

	finp := file.Open(env.AllPhraseDivisionBodyPath)
	scan := file.Scanner(finp)
	for scannabled, text := scan(); scannabled; scannabled, text = scan() {
		cName, minClusterDist := "", -1
		for clusterName, clusterTexts := range base {
			minDist := -1
			for _, cText := range clusterTexts {
				dist := levenshtein.NewSentence(cText).Distance(text)
				if minDist < 0 || dist < minDist {
					cName, minDist = clusterName, dist
				}
			}
			if minClusterDist == minDist {
				cName = "same"
			}
			if minDist < minClusterDist {
				minClusterDist = minDist
			}
		}
		base[cName] = append(base[cName], text)
	}
	file.Close(finp)

	dir := env.ClusterClassifiedDir
	for cName, cTexts := range base {
		foutp := file.CreateOpen(dir + "/" + cName + ".txt")

		for _, cText := range cTexts {
			foutp.WriteString(cText + "\n")
		}

		file.Close(foutp)
	}
}

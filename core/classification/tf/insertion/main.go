package main

import (
	"github.com/yut-kt/high-q-TTS/env"
	"github.com/yut-kt/high-q-TTS/peripheral/analysis"
	"os"
	"github.com/yut-kt/high-q-TTS/peripheral/file"
	"github.com/yut-kt/high-q-TTS/repository"
	"github.com/yut-kt/high-q-TTS/entity"
	"fmt"
)

func main() {
	finp, err := os.Open(env.AllUniqueBodyPath)
	if err != nil {
		panic(err)
	}
	defer finp.Close()

	scan := file.Scanner(finp)

	m := analysis.BuildMeCab()
	defer m.Destroy()

	dbSession, err := repository.NewSession()
	if err != nil {
		panic(err)
	}
	tfRepository := repository.TFRepository{Session: dbSession}

	scannabled, text := scan()
	for line := 1; scannabled; line++ {
		if text == "" {
			panic("contains blank line")
		}
		node := m.ParseToNode(text)
		tfMap := node.MakeTFMap()
		tfRepository.Create(&entity.TF{Line: line, Str: tfMap.Compress()})

		scannabled, text = scan()
	}
}

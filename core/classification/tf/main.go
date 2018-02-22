package main

import (
	"github.com/yut-kt/high-q-TTS/env"
	"github.com/yut-kt/high-q-TTS/peripheral/analysis"
	"os"
	"github.com/yut-kt/high-q-TTS/peripheral/file"
)

func main() {
	finp, err := os.Open(env.AllBodyPath)
	if err != nil {
		panic(err)
	}
	defer finp.Close()

	scan := file.Scanner(finp)

	m := analysis.BuildMeCab()
	defer m.Destroy()

	//dbSession, err := repository.NewSession()
	if err != nil {
		panic(err)
	}
	//tfRepository := repository.TFRepository{dbSession}

	scannabled, text, err := scan()
	if err != nil {
		panic(err)
	}
	for line := 1; scannabled; line++ {
		node := m.ParseToNode(text)
		println(node.MakeWakatiStr())
		//tfMap := node.MakeTFMap()

		//tfRepository.Create(&entity.TF{Line: line, Str: tfMap.Compress()})

		scannabled, text, err = scan()
		if err != nil {
			panic(err)
		}
	}
}

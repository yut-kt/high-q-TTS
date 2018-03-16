package main

import "github.com/yut-kt/high-q-TTS/repository"

func main()  {
	dbSession, err := repository.NewSession()
	if err != nil {
		panic(err)
	}
	tfRepository := repository.TFRepository{Session: dbSession}

	tfs, err := tfRepository.Select()
	if err != nil {
		panic(err)
	}

	for _, tf := range tfs {

	}
	println(len(a))
}

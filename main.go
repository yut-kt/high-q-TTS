package main

import "github.com/yut-kt/high-q-TTS/repository"

func main()  {
	dbSession, err := repository.NewSession()
	if err != nil {
		panic(err)
	}
	tfRepository := repository.TFRepository{Session: dbSession}

	a, err := tfRepository.Select()
	if err != nil {
		panic(err)
	}
	println(len(a))
}

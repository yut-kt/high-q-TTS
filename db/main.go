package main

import "github.com/yut-kt/high-q-TTS/db/migration"

func main() {
	migration.Down()
	migration.Up()
}

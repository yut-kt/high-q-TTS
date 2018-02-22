package repository

import "github.com/yut-kt/high-q-TTS/entity"

type WakatiRepository struct {
	*Session
}

func (repo *WakatiRepository) Create(wakati *entity.Wakati) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		// Roll back if panic happens
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if _, err := repo.db.Exec(`INSERT INTO wakati (line, str) VALUES (?, ?)`, wakati.Line, wakati.Str); err != nil {
		return err
	}

	return nil
}

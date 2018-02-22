package repository

import "github.com/yut-kt/high-q-TTS/entity"

type TFRepository struct {
	*Session
}

func (repo *TFRepository) Create(tf *entity.TF) error {
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

	if _, err := repo.db.Exec(`INSERT INTO tf (line, str) VALUES (?, ?)`, tf.Line, tf.Str); err != nil {
		return err
	}

	return nil
}

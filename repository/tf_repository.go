package repository

import (
	"github.com/yut-kt/high-q-TTS/entity"
)

type TFRepository struct {
	*Session
}

func (repo *TFRepository) Select() ([]*entity.TF, error) {
	rows, err := repo.db.Query(`SELECT * FROM tf`)
	if err != nil {
		return nil, err
	}

	tfs := make([]*entity.TF, 0)
	for rows.Next() {
		tf := entity.TF{}
		err := rows.Scan(&tf.Line, &tf.Str, &tf.Class)
		if err != nil {
			return nil, err
		}
		tfs = append(tfs, &tf)
	}
	return tfs, nil
}

func (repo *TFRepository) Create(tf *entity.TF) error {
	if _, err := repo.db.Exec(`INSERT INTO tf (line, str) VALUES (?, ?)`, tf.Line, tf.Str); err != nil {
		return err
	}
	return nil
}

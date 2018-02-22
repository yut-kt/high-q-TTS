package repository

import "database/sql"
import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/yut-kt/high-q-TTS/env"
)

type Session struct {
	db *sql.DB
}

func NewSession() (*Session, error) {
	s := new(Session)

	err := s.build()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Session) build() error {
	db, err := sql.Open(env.DataBaseDriver, env.DataBase)
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

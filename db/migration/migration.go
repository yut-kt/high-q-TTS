package migration

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yut-kt/high-q-TTS/env"
)

type Session struct {
	db *sql.DB
}

func newSession() (*Session, error) {
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

func Up() (err error) {

	session, err := newSession()
	if err != nil {
		return err
	}
	defer func() {
		err = session.db.Close()
	}()

	if _, err = session.db.Exec(`CREATE TABLE "tf" ("line" INTEGER PRIMARY KEY, "str" VARCHAR(255), "class" VARCHAR(1) )`); err != nil {
		return err
	}
	if _, err = session.db.Exec(`CREATE TABLE "wakati" ("line" INTEGER PRIMARY KEY, "str" VARCHAR(255), "class" VARCHAR(1) )`); err != nil {
		return err
	}
	if _, err = session.db.Exec(`CREATE TABLE "answer" ("wakati_str" VARCHAR(255), "class" VARCHAR(1) )`); err != nil {
		return err
	}
	return
}

func Down() (err error) {
	session, err := newSession()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = session.db.Close()
	}()

	if _, err = session.db.Exec(`DROP TABLE IF EXISTS "tf"`); err != nil {
		return err
	}
	if _, err = session.db.Exec(`DROP TABLE IF EXISTS "wakati"`); err != nil {
		return err
	}
	if _, err = session.db.Exec(`DROP TABLE IF EXISTS "answer"`); err != nil {
		return
	}
	return
}

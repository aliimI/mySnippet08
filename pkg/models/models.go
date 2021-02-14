package models

import (
	"database/sql"
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
type Example struct {
	DB *sql.DB
}

func (m *Example) ExampleTransaction() error {
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO ...")
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("UPDATE  ...")
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err

}

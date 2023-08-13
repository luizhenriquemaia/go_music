package config

import (
	"database/sql"
	"errors"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("record doesn't exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SqliteRepositoryAbstract interface {
	All()
	GetByID()
	Create()
	Update()
	Delete()
}

type SqliteRepository struct {
	db *sql.DB
}

func NewSqliteRepository(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (r *SqliteRepository) Migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS album(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(50) NOT NULL,
			artist VARCHAR(50) NOT NULL,
			price DECIMAL(10,2) DEFAULT '0.00'
		);
	`
	_, err := r.db.Exec(query)
	return err
}
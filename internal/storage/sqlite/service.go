package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // SQL database itself, with database/sql
)

// Service is a SQLite storage service
type Service struct {
	db *sql.DB
}

// New returns a SQLite storage service
func New(path string) (*Service, error) {
	dsn := fmt.Sprintf("file:%s?_foreign_keys=true&cache=shared", path)
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return &Service{db}, nil
}

// Close closes the database
func (s *Service) Close() error {
	return s.db.Close()
}

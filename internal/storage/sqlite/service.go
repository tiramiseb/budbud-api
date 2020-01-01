package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQL database itself, with database/sql
)

// Service is a SQLite storage service
type Service struct {
	db *sql.DB
}

// New returns a SQLite storage service
func New(path string) (*Service, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	s := &Service{db}
	return s, s.checkInit()
}

// Close closes the database
func (s *Service) Close() error {
	return s.db.Close()
}

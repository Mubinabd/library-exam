package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Mubinabd/library-exam/config"
	"github.com/Mubinabd/library-exam/storage"
	_ "github.com/lib/pq"
)

type Storage struct {
	db        *sql.DB
	BookS     storage.BookI
	BorrowerS storage.BorrowerI
	GenreS    storage.GenreI
	AuthorS   storage.AuthorI
}

func ConnectDB() (*Storage, error) {
	cfg := config.Load()
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		"postgres-db",
		cfg.PostgresPort,
		cfg.PostgresDatabase)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	authS := NewAuthorStorge(db)
	bookS := NewBookStorage(db)
	borrowerS := NewBorrowerStrorage(db)
	genreS := NewGenreStorage(db)
	return &Storage{
		db:        db,
		AuthorS:   authS,
		BookS:     bookS,
		BorrowerS: borrowerS,
		GenreS:    genreS,
	}, nil
}
func (s *Storage) Author() storage.AuthorI {
	if s.AuthorS == nil {
		s.AuthorS = NewAuthorStorge(s.db)
	}
	return s.AuthorS
}

func (s *Storage) Book() storage.BookI {
	if s.BookS == nil {
		s.BookS = NewBookStorage(s.db)
	}
	return s.BookS
}

func (s *Storage) Borrower() storage.BorrowerI {
	if s.BorrowerS == nil {
		s.BorrowerS = NewBorrowerStrorage(s.db)
	}
	return s.BorrowerS
}

func (s *Storage) Genre() storage.GenreI {
	if s.GenreS == nil {
		s.GenreS = NewGenreStorage(s.db)
	}
	return s.GenreS
}

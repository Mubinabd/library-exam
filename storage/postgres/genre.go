package postgres

import (
	"database/sql"
	"fmt"
	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/google/uuid"
	"strings"
)

type GenreStorage struct {
	db *sql.DB
}

func NewGenreStorage(db *sql.DB) *GenreStorage {
	return &GenreStorage{db: db}
}

func (genre *GenreStorage) CreateGenre(req *pb.GenreCreate) (*pb.Genre, error) {
	id := uuid.NewString()
	req.Id = id
	query := `
		INSERT INTO genre
		(id, name)
		VALUES($1, $2)
		RETURNING id, name`
	var resp pb.Genre
	err := genre.db.QueryRow(query, req.Id, req.Name).
		Scan(&resp.Id, &resp.Name)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (genre *GenreStorage) GetGenre(req *pb.ByName) (*pb.GenreCreate, error) {
	query := `
		SELECT id, name
		FROM genre
		WHERE name = $1`
	var resp pb.GenreCreate
	err := genre.db.QueryRow(query, req.Name).Scan(&resp.Id, &resp.Name)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (genre *GenreStorage) GetAllGenres(req *pb.NameFilter) (*pb.Genres, error) {
	query := `
		SELECT 
		id,name
		FROM genre
		WHERE deleted_at = 0`
	var conditions []string
	var args []interface{}
	paramIndex := 1

	if req.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name = $%d", paramIndex))
		args = append(args,req.Name)
		paramIndex++
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := genre.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genres := pb.Genres{}
	for rows.Next() {
		genre := pb.Genre{}
		err := rows.Scan(&genre.Id,&genre.Name)
		if err != nil {
			return nil, err
		}
		genres.Genres = append(genres.Genres, &genre)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &genres, nil
}

func (genre *GenreStorage) UpdateGenre(req *pb.GenreCreate) (*pb.Void, error) {
	query := `UPDATE genre SET
		name = $1,
		updated_at = now()
		WHERE id = $2 AND deleted_at = 0`
	_, err := genre.db.Exec(query, req.Name, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (genre *GenreStorage) DeleteGenre(req *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE genre 
		SET deleted_at = EXTRACT(EPOCH FROM NOW()) 
		WHERE id = $1`
	_, err := genre.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (genre *GenreStorage)GetBooksByGenre(req *pb.GenreId) (*pb.GenreBooks, error) {
	query := `
		SELECT 
			id, title, author_id, genre_id, summary 
		FROM 
			book
		WHERE 
			deleted_at = 0 AND genre_id = $1`

	rows, err := genre.db.Query(query, req.GenreId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := pb.GenreBooks{}

	for rows.Next() {
		book := pb.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.AuthorId, &book.GenreId, &book.Summary)
		if err != nil {
			return nil, err
		}

		books.Books = append(books.Books, &book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &books, nil
}

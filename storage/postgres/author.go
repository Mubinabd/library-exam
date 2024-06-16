package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/google/uuid"
)

type AuthorStorage struct {
	db *sql.DB
}

func NewAuthorStorge(db *sql.DB) *AuthorStorage {
	return &AuthorStorage{db: db}
}

func (auth *AuthorStorage) CreateAuthor(req *pb.AuthorCreate) (*pb.Author, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO author
		(id, name, biography)
		VALUES($1, $2, $3)
		RETURNING id, name, biography`

	var resp pb.Author
	err := auth.db.QueryRow(query, id, req.Name, req.Biography).
		Scan(&resp.Id, &resp.Name, &resp.Biography)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (auth *AuthorStorage) GetAuthor(req *pb.ById) (*pb.Author, error) {
	query := `
		SELECT id, name, biography
		FROM author
		WHERE id = $1`
	var resp pb.Author

	err := auth.db.QueryRow(query, req.Id).Scan(&resp.Id, &resp.Name, &resp.Biography)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (auth *AuthorStorage) UpdateAuthor(req *pb.AuthorCreate) (*pb.Void, error) {
	query := `UPDATE author SET
	name = $1,
	biography = $2,
	updated_at = now()
	WHERE id = $3 AND deleted_at = 0`
	_, err := auth.db.Exec(query, req.Name, req.Biography, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (auth *AuthorStorage) DeleteAuthor(req *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE author 
		SET deleted_at = EXTRACT(EPOCH FROM NOW()) 
		WHERE id = $1`
	_, err := auth.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (s *AuthorStorage) GetAllAuthors(req *pb.NameFilter) (*pb.Authors, error) {
	query := `
		SELECT 
			id, name, biography 
		FROM author
		WHERE deleted_at = 0`

	var conditions []string
	var args []interface{}
	paramIndex := 1

	if req.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name = $%d", paramIndex))
		args = append(args, req.Name)
		paramIndex++
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors pb.Authors
	for rows.Next() {
		var author pb.Author
		err := rows.Scan(&author.Id, &author.Name, &author.Biography)
		if err != nil {
			return nil, err
		}
		authors.Authors = append(authors.Authors, &author)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &authors, nil
}

func (s *AuthorStorage)GetAuthorBooks(req *pb.AuthorID) (*pb.UserBook, error) {
	query := `
		SELECT 
			id, title, author_id, genre_id, summary 
		FROM 
			book
		WHERE 
			deleted_at = 0 AND author_id = $1`

	rows, err := s.db.Query(query, req.AuthorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := pb.UserBook{}

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

package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/google/uuid"
)

type BookStorage struct {
	db *sql.DB
}

func NewBookStorage(db *sql.DB) *BookStorage {
	return &BookStorage{db: db}
}

func (book *BookStorage) CreateBook(req *pb.BookCreate) (*pb.Book, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO book (id, title, author_id, genre_id, summary)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, author_id, genre_id, summary`

	var resp pb.Book
	err := book.db.QueryRow(query, id, req.Title, req.AuthorID, req.GenreID, req.Summary).
		Scan(&resp.Id, &resp.Title, &resp.AuthorId, &resp.GenreId, &resp.Summary)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (book *BookStorage) GetBook(req *pb.ByTitle) (*pb.Book, error) {
	query := `
		SELECT id, title, author_id, genre_id, summary
		FROM book
		WHERE title = $1`
	var resp pb.Book
	err := book.db.QueryRow(query, req.Title).Scan(&resp.Id, &resp.Title, &resp.AuthorId, &resp.GenreId, &resp.Summary)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (book *BookStorage) UpdateBook(req *pb.BookCreate) (*pb.Void, error) {
	query := `UPDATE book SET
	title = $1,
	author_id = $2,
	genre_id = $3,
	summary = $4,
	updated_at = now()
	WHERE id = $5 AND deleted_at = 0`
	_, err := book.db.Exec(query, req.Title, req.AuthorID, req.GenreID, req.Summary, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (book *BookStorage) DeleteBook(req *pb.ById) (*pb.Void, error) {

	query := `
		UPDATE book 
		SET deleted_at = EXTRACT(EPOCH FROM NOW()) 
		WHERE id = $1`
	_, err := book.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (book *BookStorage) GetAllBooks(req *pb.TitleFilter) (*pb.Books, error) {
	query := `
		SELECT 
		id,title, author_id, genre_id, summary 
		FROM book
		WHERE deleted_at = 0`

	var conditions []string
	var args []interface{}
	paramIndex := 1

	if req.Title != "" {
		conditions = append(conditions, fmt.Sprintf("title = $%d", paramIndex))
		args = append(args, req.Title)
		paramIndex++
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := book.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := pb.Books{}
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

func (book *BookStorage) SearchTitleAndAuthor(req *pb.Search) (*pb.Books, error) {

	query := `
		SELECT 
		id,title, author_id, genre_id, summary 
		FROM book
		WHERE deleted_at = 0 AND title LIKE $1 OR summary LIKE $2`
	rows, err := book.db.Query(query, "%"+req.Title+"%", "%"+req.Author+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := pb.Books{}
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

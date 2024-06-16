package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/Mubinabd/library-exam/storage/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	bookStorage := postgres.NewBookStorage(db)

	req := &pb.BookCreate{
		Title:    "Test Book",
		AuthorID: "author-1",
		GenreID:  "genre-1",
		Summary:  "Test Summary",
	}

	mock.ExpectQuery("INSERT INTO book").
		WithArgs(sqlmock.AnyArg(), req.Title, req.AuthorID, req.GenreID, req.Summary).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author_id", "genre_id", "summary"}).AddRow("1", req.Title, req.AuthorID, req.GenreID, req.Summary))

	resp, err := bookStorage.CreateBook(req)
	require.NoError(t, err)

	assert.Equal(t, "1", resp.Id)
	assert.Equal(t, req.Title, resp.Title)
	assert.Equal(t, req.AuthorID, resp.AuthorId)
	assert.Equal(t, req.GenreID, resp.GenreId)
	assert.Equal(t, req.Summary, resp.Summary)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	bookStorage := postgres.NewBookStorage(db)

	req := &pb.ByTitle{Title: "Test Book"}

	mock.ExpectQuery("SELECT id, title, author_id, genre_id, summary FROM book WHERE title = \\$1").
		WithArgs(req.Title).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author_id", "genre_id", "summary"}).AddRow("1", req.Title, "author-1", "genre-1", "Test Summary"))

	resp, err := bookStorage.GetBook(req)
	require.NoError(t, err)

	assert.Equal(t, "1", resp.Id)
	assert.Equal(t, req.Title, resp.Title)
	assert.Equal(t, "author-1", resp.AuthorId)
	assert.Equal(t, "genre-1", resp.GenreId)
	assert.Equal(t, "Test Summary", resp.Summary)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	bookStorage := postgres.NewBookStorage(db)

	req := &pb.BookCreate{
		Id:       "1",
		Title:    "Updated Book",
		AuthorID: "author-1",
		GenreID:  "genre-1",
		Summary:  "Updated Summary",
	}

	mock.ExpectExec("UPDATE book SET").
		WithArgs(req.Title, req.AuthorID, req.GenreID, req.Summary, req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = bookStorage.UpdateBook(req)
	require.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	bookStorage := postgres.NewBookStorage(db)

	req := &pb.ById{Id: "1"}

	mock.ExpectExec("UPDATE book SET deleted_at").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = bookStorage.DeleteBook(req)
	require.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	bookStorage := postgres.NewBookStorage(db)

	req := &pb.TitleFilter{Title: ""}

	mock.ExpectQuery("SELECT id,title, author_id, genre_id, summary FROM book WHERE deleted_at = 0").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author_id", "genre_id", "summary"}).
			AddRow("1", "Book1", "author-1", "genre-1", "Summary1").
			AddRow("2", "Book2", "author-2", "genre-2", "Summary2"))

	resp, err := bookStorage.GetAllBooks(req)
	require.NoError(t, err)

	require.Len(t, resp.Books, 2)
	assert.Equal(t, "1", resp.Books[0].Id)
	assert.Equal(t, "Book1", resp.Books[0].Title)
	assert.Equal(t, "author-1", resp.Books[0].AuthorId)
	assert.Equal(t, "genre-1", resp.Books[0].GenreId)
	assert.Equal(t, "Summary1", resp.Books[0].Summary)

	assert.Equal(t, "2", resp.Books[1].Id)
	assert.Equal(t, "Book2", resp.Books[1].Title)
	assert.Equal(t, "author-2", resp.Books[1].AuthorId)
	assert.Equal(t, "genre-2", resp.Books[1].GenreId)
	assert.Equal(t, "Summary2", resp.Books[1].Summary)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSearchTitleAndAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	bookStorage := postgres.NewBookStorage(db)

	req := &pb.Search{Title: "Test"}

	mock.ExpectQuery("SELECT id,title, author_id, genre_id, summary FROM book WHERE deleted_at = 0 AND title LIKE \\$1 OR summary LIKE \\$2").
		WithArgs("%"+req.Title+"%", "%"+req.Title+"%").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author_id", "genre_id", "summary"}).
			AddRow("1", "Test Book", "author-1", "genre-1", "Test Summary"))

	resp, err := bookStorage.SearchTitleAndAuthor(req)
	require.NoError(t, err)

	require.Len(t, resp.Books, 1)
	assert.Equal(t, "1", resp.Books[0].Id)
	assert.Equal(t, "Test Book", resp.Books[0].Title)
	assert.Equal(t, "author-1", resp.Books[0].AuthorId)
	assert.Equal(t, "genre-1", resp.Books[0].GenreId)
	assert.Equal(t, "Test Summary", resp.Books[0].Summary)

	assert.NoError(t, mock.ExpectationsWereMet())
}

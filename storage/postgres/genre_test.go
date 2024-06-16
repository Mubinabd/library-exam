package postgres_test

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/Mubinabd/library-exam/storage/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateGenre(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	genreStorage := postgres.NewGenreStorage(db)

	req := &pb.GenreCreate{
		Name: "Test Genre",
	}

	mock.ExpectQuery("INSERT INTO genre").
		WithArgs(sqlmock.AnyArg(), req.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("1", req.Name))

	resp, err := genreStorage.CreateGenre(req)
	require.NoError(t, err)

	assert.Equal(t, "1", resp.Id)
	assert.Equal(t, req.Name, resp.Name)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetGenre(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	genreStorage := postgres.NewGenreStorage(db)

	req := &pb.ByName{Name: "Test Genre"}

	mock.ExpectQuery("SELECT id, name FROM genre WHERE name = \\$1").
		WithArgs(req.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("1", req.Name))

	resp, err := genreStorage.GetGenre(req)
	require.NoError(t, err)

	assert.Equal(t, "1", resp.Id)
	assert.Equal(t, req.Name, resp.Name)

	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestUpdateGenre(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	genreStorage := postgres.NewGenreStorage(db)

	req := &pb.GenreCreate{
		Id:   "1",
		Name: "Updated Genre",
	}

	mock.ExpectExec("UPDATE genre SET").
		WithArgs(req.Name, req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = genreStorage.UpdateGenre(req)
	require.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteGenre(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	genreStorage := postgres.NewGenreStorage(db)

	req := &pb.ById{Id: "1"}

	mock.ExpectExec("UPDATE genre SET deleted_at").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = genreStorage.DeleteGenre(req)
	require.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetBooksByGenre(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	genreStorage := postgres.NewGenreStorage(db)

	req := &pb.GenreId{GenreId: "1"}

	mock.ExpectQuery("SELECT id, title, author_id, genre_id, summary FROM book WHERE deleted_at = 0 AND genre_id = \\$1").
		WithArgs(req.GenreId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author_id", "genre_id", "summary"}).
			AddRow("1", "Test Book 1", "author-1", req.GenreId, "Summary 1").
			AddRow("2", "Test Book 2", "author-2", req.GenreId, "Summary 2"))

	resp, err := genreStorage.GetBooksByGenre(req)
	require.NoError(t, err)

	require.Len(t, resp.Books, 2)
	assert.Equal(t, "1", resp.Books[0].Id)
	assert.Equal(t, "Test Book 1", resp.Books[0].Title)
	assert.Equal(t, "author-1", resp.Books[0].AuthorId)
	assert.Equal(t, req.GenreId, resp.Books[0].GenreId)
	assert.Equal(t, "Summary 1", resp.Books[0].Summary)

	assert.Equal(t, "2", resp.Books[1].Id)
	assert.Equal(t, "Test Book 2", resp.Books[1].Title)
	assert.Equal(t, "author-2", resp.Books[1].AuthorId)
	assert.Equal(t, req.GenreId, resp.Books[1].GenreId)
	assert.Equal(t, "Summary 2", resp.Books[1].Summary)

	assert.NoError(t, mock.ExpectationsWereMet())
}

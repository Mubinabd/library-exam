package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mubinabd/library-exam/genproto"
	"github.com/Mubinabd/library-exam/storage/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	authorStorage := postgres.NewAuthorStorge(db)

	req := &genproto.AuthorCreate{
		Name:      "Test Author",
		Biography: "Test Biography",
	}

	mock.ExpectQuery("INSERT INTO author").
		WithArgs(sqlmock.AnyArg(), req.Name, req.Biography).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "biography"}).AddRow("c2cd0d50-c04b-4eef-9a59-5e253b2702e1", req.Name, req.Biography))

	resp, err := authorStorage.CreateAuthor(req)
	require.NoError(t, err)

	assert.Equal(t, "c2cd0d50-c04b-4eef-9a59-5e253b2702e1", resp.Id)
	assert.Equal(t, req.Name, resp.Name)
	assert.Equal(t, req.Biography, resp.Biography)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	authorStorage := postgres.NewAuthorStorge(db)

	req := &genproto.ById{Id: "1"}

	mock.ExpectQuery("SELECT id, name, biography FROM author WHERE id = \\$1").
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "biography"}).AddRow("c2cd0d50-c04b-4eef-9a59-5e253b2702e1", "Test Author", "Test Biography"))

	resp, err := authorStorage.GetAuthor(req)
	require.NoError(t, err)

	assert.Equal(t, "c2cd0d50-c04b-4eef-9a59-5e253b2702e1", resp.Id)
	assert.Equal(t, "Test Author", resp.Name)
	assert.Equal(t, "Test Biography", resp.Biography)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	authorStorage := postgres.NewAuthorStorge(db)

	req := &genproto.AuthorCreate{
		Id:        "c2cd0d50-c04b-4eef-9a59-5e253b2702e1",
		Name:      "Updated Author",
		Biography: "Updated Biography",
	}

	mock.ExpectExec("UPDATE author SET").
		WithArgs(req.Name, req.Biography, req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = authorStorage.UpdateAuthor(req)
	require.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	authorStorage := postgres.NewAuthorStorge(db)

	req := &genproto.ById{Id: "c2cd0d50-c04b-4eef-9a59-5e253b2702e1"}

	mock.ExpectExec("UPDATE author SET deleted_at").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = authorStorage.DeleteAuthor(req)
	require.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllAuthors(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	authorStorage := postgres.NewAuthorStorge(db)

	req := &genproto.NameFilter{Name: ""}

	mock.ExpectQuery("SELECT id, name, biography FROM author WHERE deleted_at = 0").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "biography"}).
			AddRow("c2cd0d50-c04b-4eef-9a59-5e253b2702e1", "Author1", "Bio1").
			AddRow("3f7d1529-8139-43b8-8f28-39f4d0401866", "Author2", "Bio2"))

	resp, err := authorStorage.GetAllAuthors(req)
	require.NoError(t, err)

	require.Len(t, resp.Authors, 2)
	assert.Equal(t, "c2cd0d50-c04b-4eef-9a59-5e253b2702e1", resp.Authors[0].Id)
	assert.Equal(t, "Author1", resp.Authors[0].Name)
	assert.Equal(t, "Bio1", resp.Authors[0].Biography)

	assert.Equal(t, "3f7d1529-8139-43b8-8f28-39f4d0401866", resp.Authors[1].Id)
	assert.Equal(t, "Author2", resp.Authors[1].Name)
	assert.Equal(t, "Bio2", resp.Authors[1].Biography)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAuthorBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	authorStorage := postgres.NewAuthorStorge(db)

	req := &genproto.AuthorID{AuthorId: "c2cd0d50-c04b-4eef-9a59-5e253b2702e1"}

	mock.ExpectQuery("SELECT id, title, author_id, genre_id, summary FROM book WHERE deleted_at = 0 AND author_id = \\$1").
		WithArgs(req.AuthorId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author_id", "genre_id", "summary"}).
			AddRow("c2cd0d50-c04b-4eef-9a59-5e253b2702e1", "Book1", "1", "1", "Summary1").
			AddRow("3f7d1529-8139-43b8-8f28-39f4d0401866", "Book2", "1", "1", "Summary2"))

	resp, err := authorStorage.GetAuthorBooks(req)
	require.NoError(t, err)

	require.Len(t, resp.Books, 2)
	assert.Equal(t, "c2cd0d50-c04b-4eef-9a59-5e253b2702e1", resp.Books[0].Id)
	assert.Equal(t, "Book1", resp.Books[0].Title)
	assert.Equal(t, "1", resp.Books[0].AuthorId)
	assert.Equal(t, "1", resp.Books[0].GenreId)
	assert.Equal(t, "Summary1", resp.Books[0].Summary)

	assert.Equal(t, "3f7d1529-8139-43b8-8f28-39f4d0401866", resp.Books[1].Id)
	assert.Equal(t, "Book2", resp.Books[1].Title)
	assert.Equal(t, "1", resp.Books[1].AuthorId)
	assert.Equal(t, "1", resp.Books[1].GenreId)
	assert.Equal(t, "Summary2", resp.Books[1].Summary)

	assert.NoError(t, mock.ExpectationsWereMet())
}

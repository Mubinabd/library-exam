package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/Mubinabd/library-exam/storage/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateBorrower(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	borrowerService := postgres.NewBorrowerStrorage(db)

	req := &pb.BorrowerCreate{
		UserID:     "user-1",
		BookID:     "book-1",
		BorrowDate: "2024-01-01",
		ReturnDate: "2024-01-10",
	}

	mock.ExpectQuery("INSERT INTO borrower").
		WithArgs(sqlmock.AnyArg(), req.UserID, req.BookID, req.BorrowDate, req.ReturnDate).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "book_id", "borrow_date", "return_date"}).AddRow("1", req.UserID, req.BookID, req.BorrowDate, req.ReturnDate))

	resp, err := borrowerService.CreateBorrower(req)
	require.NoError(t, err)

	assert.Equal(t, "1", resp.Id)
	assert.Equal(t, req.UserID, resp.UserID)
	assert.Equal(t, req.BookID, resp.BookID)
	assert.Equal(t, req.BorrowDate, resp.BorrowDate)
	assert.Equal(t, req.ReturnDate, resp.ReturnDate)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetBorrower(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	borrowerService := postgres.NewBorrowerStrorage(db)

	req := &pb.ById{Id: "1"}

	mock.ExpectQuery("SELECT id, user_id, book_id, borrow_date, return_date FROM borrower WHERE id = \\$1").
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "book_id", "borrow_date", "return_date"}).AddRow("1", "user-1", "book-1", "2024-01-01", "2024-01-10"))

	resp, err := borrowerService.GetBorrower(req)
	require.NoError(t, err)

	assert.Equal(t, "1", resp.Id)
	assert.Equal(t, "user-1", resp.UserID)
	assert.Equal(t, "book-1", resp.BookID)
	assert.Equal(t, "2024-01-01", resp.BorrowDate)
	assert.Equal(t, "2024-01-10", resp.ReturnDate)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllBorrowers(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	borrowerService := postgres.NewBorrowerStrorage(db)

	mock.ExpectQuery("SELECT id,user_id, book_id, borrow_date, return_date FROM borrower WHERE deleted_at = 0").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "book_id", "borrow_date", "return_date"}).
			AddRow("6f92a86a-efcf-4aa6-8b31-705ba632976f", "user-1", "book-1", "2024-01-01", "2024-01-10").
			AddRow("eb88da3a-4fd9-4e57-9ec5-6142bb66b1d7", "user-2", "book-2", "2024-01-02", "2024-01-11"))

	resp, err := borrowerService.GetAllBorrowers(&pb.Void{})
	require.NoError(t, err)

	require.Len(t, resp.Borrowers, 2)

	assert.Equal(t, "6f92a86a-efcf-4aa6-8b31-705ba632976f", resp.Borrowers[0].Id)
	assert.Equal(t, "user-1", resp.Borrowers[0].UserID)
	assert.Equal(t, "book-1", resp.Borrowers[0].BookID)
	assert.Equal(t, "2024-01-01", resp.Borrowers[0].BorrowDate)
	assert.Equal(t, "2024-01-10", resp.Borrowers[0].ReturnDate)

	assert.Equal(t, "eb88da3a-4fd9-4e57-9ec5-6142bb66b1d7", resp.Borrowers[1].Id)
	assert.Equal(t, "user-2", resp.Borrowers[1].UserID)
	assert.Equal(t, "book-2", resp.Borrowers[1].BookID)
	assert.Equal(t, "2024-01-02", resp.Borrowers[1].BorrowDate)
	assert.Equal(t, "2024-01-11", resp.Borrowers[1].ReturnDate)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateBorrower(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	borrowerService := postgres.NewBorrowerStrorage(db)

	req := &pb.BorrowerCreate{
		Id:         "1",
		UserID:     "user-1",
		BookID:     "book-1",
		BorrowDate: "2024-01-01",
		ReturnDate: "2024-01-10",
	}

	mock.ExpectExec("UPDATE borrower SET").
		WithArgs(req.UserID, req.BookID, req.BorrowDate, req.ReturnDate, req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = borrowerService.UpdateBorrower(req)
	require.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteBorrower(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	borrowerService := postgres.NewBorrowerStrorage(db)

	req := &pb.ById{Id: "1"}

	mock.ExpectExec("UPDATE borrower SET deleted_at").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = borrowerService.DeleteBorrower(req)
	require.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

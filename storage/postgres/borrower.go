package postgres

import (
	"database/sql"
	"time"

	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/google/uuid"
)

type BorrowerService struct {
	db *sql.DB
}

func NewBorrowerStrorage(db *sql.DB) *BorrowerService {
	return &BorrowerService{db: db}
}

func (borrower *BorrowerService) CreateBorrower(req *pb.BorrowerCreate) (*pb.Borrower, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO borrower
		(id, user_id, book_id, borrow_date, return_date)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id,user_id, book_id, borrow_date, return_date`
	var resp pb.Borrower
	err := borrower.db.QueryRow(query, id, req.UserID, req.BookID, req.BorrowDate, req.ReturnDate).
		Scan(&resp.Id, &resp.UserID, &resp.BookID, &resp.BorrowDate, &resp.ReturnDate)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (borrower *BorrowerService) GetBorrower(req *pb.ById) (*pb.Borrower, error) {
	query := `
		SELECT id, user_id, book_id, borrow_date, return_date
		FROM borrower
		WHERE id = $1`
	var resp pb.Borrower
	err := borrower.db.QueryRow(query, req.Id).Scan(&resp.Id, &resp.UserID, &resp.BookID, &resp.BorrowDate, &resp.ReturnDate)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (borrower *BorrowerService) GetAllBorrowers(req *pb.Void) (*pb.Borrowers, error) {
	query := `
		SELECT 
		user_id, book_id, borrow_date, return_date 
		FROM borrower`
	rows, err := borrower.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var borrowers pb.Borrowers
	for rows.Next() {
		var borrower pb.Borrower
		if err := rows.Scan(&borrower.UserID, &borrower.BookID, &borrower.BorrowDate, &borrower.ReturnDate); err != nil {
			return nil, err
		}
		borrowers.Borrowers = append(borrowers.Borrowers, &borrower)
	}
	return &borrowers, nil
}

func (borrower *BorrowerService) UpdateBorrower(req *pb.BorrowerCreate) (*pb.Void, error) {
	query := `UPDATE borrower SET
			user_id = $1,
			book_id = $2,
			borrow_date = $3,
			return_date = $4,
			updated_at = now()
			WHERE id = $5 AND deleted_at = 0`
	_, err := borrower.db.Exec(query, req.UserID, req.BookID, req.BorrowDate, req.ReturnDate, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (borrower *BorrowerService) DeleteBorrower(req *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE borrower 
		SET deleted_at = EXTRACT(EPOCH FROM NOW()) 
		WHERE id = $1`
	_, err := borrower.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (borrower *BorrowerService) BorrowerBooks(req *pb.UserId) (*pb.BorrowedBooks, error) {
	query := `
        SELECT 
            book.id, book.title, book.author_id, book.genre_id, book.summary, 
            borrower.borrow_date, borrower.return_date
        FROM 
            book
        JOIN
            borrower ON book.id = borrower.book_id
        WHERE 
            book.deleted_at = 0 AND borrower.user_id = $1`

	rows, err := borrower.db.Query(query, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var borrowedBooks pb.BorrowedBooks

	for rows.Next() {
		book := pb.Book{}
		borrowedB := pb.BorrowedBook{}
		err := rows.Scan(&book.Id, &book.Title, &book.AuthorId, &book.GenreId, &book.Summary, &borrowedB.BorrowDate, &borrowedB.ReturnDate)
		if err != nil {
			return nil, err
		}
		borrowedB.Book = &book
		borrowedBooks.Books = append(borrowedBooks.Books, &borrowedB)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &borrowedBooks, nil
}
func(borrower *BorrowerService)GetOverdueBooks(req *pb.OverdueRequest) (*pb.BorrowedBooks, error) {
	currentDate := time.Now().Format("2006-01-02")
	query := `	
		SELECT 
			book.id, book.title, book.author_id, book.genre_id, book.summary, 
			borrower.borrow_date, borrower.return_date
		FROM 
			book
		JOIN
			borrower ON book.id = borrower.book_id
		WHERE 
			book.deleted_at = 0 AND borrower.return_date < $1`

	rows, err := borrower.db.Query(query, currentDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var borrowedBooks pb.BorrowedBooks

	for rows.Next() {
		book := pb.Book{}
		borrowedB := pb.BorrowedBook{}
		err := rows.Scan(&book.Id, &book.Title, &book.AuthorId, &book.GenreId, &book.Summary, &borrowedB.BorrowDate, &borrowedB.ReturnDate)
		if err != nil {
			return nil, err
		}
		borrowedB.Book = &book
		borrowedBooks.Books = append(borrowedBooks.Books, &borrowedB)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &borrowedBooks, nil
}



func (borrower *BorrowerService) HistoryUser(req *pb.UserId) (*pb.BorrowingHistory, error) {

	query := `
        SELECT 
            book.id, book.title, book.author_id, book.genre_id, book.summary, 
            borrower.borrow_date, borrower.return_date
        FROM 
            book
        JOIN
            borrower ON book.id = borrower.book_id
        WHERE 
           borrower.user_id = $1`

	rows, err := borrower.db.Query(query, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var borrowedBooks pb.BorrowingHistory

	for rows.Next() {
		book := pb.Book{}
		borrowedB := pb.BorrowedBook{}
		err := rows.Scan(&book.Id, &book.Title, &book.AuthorId, &book.GenreId, &book.Summary, &borrowedB.BorrowDate, &borrowedB.ReturnDate)
		if err != nil {
			return nil, err
		}
		borrowedB.Book = &book
		borrowedBooks.Books = append(borrowedBooks.Books, &borrowedB)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &borrowedBooks, nil
}

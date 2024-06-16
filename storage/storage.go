package storage

import pb "github.com/Mubinabd/library-exam/genproto"

type StorageI interface {
	Book() BookI
	Genre() GenreI
	Author() AuthorI
	Borrower() BorrowerI
}

type BookI interface {
	CreateBook(req *pb.BookCreate) (*pb.Book, error)
	UpdateBook(req *pb.BookCreate) (*pb.Void, error)
	DeleteBook(id *pb.ById) (*pb.Void, error)
	GetBook(req *pb.ByTitle) (*pb.Book, error)
	GetAllBooks(req *pb.TitleFilter) (*pb.Books, error)
	SearchTitleAndAuthor(req *pb.Search) (*pb.Books, error)
}

type GenreI interface {
	CreateGenre(req *pb.GenreCreate) (*pb.Genre, error)
	UpdateGenre(req *pb.GenreCreate) (*pb.Void, error)
	DeleteGenre(id *pb.ById) (*pb.Void, error)
	GetGenre(id *pb.ByName) (*pb.GenreCreate, error)
	GetAllGenres(req *pb.NameFilter) (*pb.Genres, error)
	GetBooksByGenre(req *pb.GenreId) (*pb.GenreBooks, error)
}
type AuthorI interface {
	CreateAuthor(req *pb.AuthorCreate) (*pb.Author, error)
	UpdateAuthor(req *pb.AuthorCreate) (*pb.Void, error)
	DeleteAuthor(id *pb.ById) (*pb.Void, error)
	GetAuthor(req *pb.ById) (*pb.Author, error)
	GetAllAuthors(req *pb.NameFilter) (*pb.Authors, error)
	GetAuthorBooks(req *pb.AuthorID) (*pb.UserBook, error)
}

type BorrowerI interface {
	CreateBorrower(req *pb.BorrowerCreate) (*pb.Borrower, error)
	UpdateBorrower(req *pb.BorrowerCreate) (*pb.Void, error)
	DeleteBorrower(id *pb.ById) (*pb.Void, error)
	GetBorrower(id *pb.ById) (*pb.Borrower, error)
	GetAllBorrowers(req *pb.Void) (*pb.Borrowers, error)
	BorrowerBooks(id *pb.UserId) (*pb.BorrowedBooks, error)
	GetOverdueBooks(req *pb.OverdueRequest) (*pb.BorrowedBooks, error)
	HistoryUser(id *pb.UserId) (*pb.BorrowingHistory, error)
}

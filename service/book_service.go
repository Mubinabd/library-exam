package service

import (
	"context"

	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/Mubinabd/library-exam/storage"
	"github.com/google/uuid"
)

type BookService struct {
	storage storage.StorageI
	pb.UnimplementedBookServiceServer
}

func NewBookStorage(storage storage.StorageI) *BookService {
	return &BookService{
		storage: storage,
	}
}

func (s *BookService)CreateBook(c context.Context, req *pb.BookCreate) (*pb.Book, error) {
	id := uuid.NewString()
	req.Id = id
	return s.storage.Book().CreateBook(req)
}
func (s *BookService)UpdateBook(c context.Context, req *pb.BookCreate) (*pb.Void, error) {
	return s.storage.Book().UpdateBook(req)
}
func (s *BookService)DeleteBook(c context.Context,id *pb.ById) (*pb.Void, error) {
	return s.storage.Book().DeleteBook(id)
}

func (s *BookService)GetBook(c context.Context,req *pb.ByTitle) (*pb.Book, error) {
	return s.storage.Book().GetBook(req)
}

func (s *BookService)GetAllBooks(c context.Context, req *pb.TitleFilter) (*pb.Books, error) {
	return s.storage.Book().GetAllBooks(req)
}

func(s *BookService)SearchTitleAndAuthor(c context.Context, req *pb.Search) (*pb.Books, error) {
	return s.storage.Book().SearchTitleAndAuthor(req)
}
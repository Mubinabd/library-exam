package service

import (
	"context"

	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/Mubinabd/library-exam/storage"
	"github.com/google/uuid"
)

type AuthorService struct {
	storage storage.StorageI
	pb.UnimplementedAuthorServiceServer
}

func NewAuthorStorage(storage storage.StorageI) *AuthorService {
	return &AuthorService{
		storage: storage,
	}
}

func (s *AuthorService) CreateAuthor(c context.Context, req *pb.AuthorCreate) (*pb.Author, error) {
	id := uuid.NewString()
	req.Id = id
	return s.storage.Author().CreateAuthor(req)
}
func (s *AuthorService) UpdateAuthor(c context.Context, req *pb.AuthorCreate) (*pb.Void, error) {
	return s.storage.Author().UpdateAuthor(req)
}
func (s *AuthorService) DeleteAuthor(c context.Context, id *pb.ById) (*pb.Void, error) {
	return s.storage.Author().DeleteAuthor(id)
}

func (s *AuthorService) GetAuthor(c context.Context, id *pb.ById) (*pb.Author, error) {
	return s.storage.Author().GetAuthor(id)
}

func (s *AuthorService) GetAllAuthors(c context.Context, req *pb.NameFilter) (*pb.Authors, error) {
	return s.storage.Author().GetAllAuthors(req)
}

func(s *AuthorService)GetAuthorBooks(c context.Context, req *pb.AuthorID) (*pb.UserBook, error) {
	return s.storage.Author().GetAuthorBooks(req)
}

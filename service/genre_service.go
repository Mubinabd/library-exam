package service

import (
	"context"

	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/Mubinabd/library-exam/storage"
	"github.com/google/uuid"
)

type GenreService struct {
	storage storage.StorageI
	pb.UnimplementedGenreServiceServer
}

func NewGenreStorage(storage storage.StorageI) *GenreService {
	return &GenreService{
		storage: storage,
	}
}

func (s *GenreService)CreateGenre(c context.Context, req *pb.GenreCreate) (*pb.Genre, error) {
	id := uuid.NewString()
	req.Id = id
	return s.storage.Genre().CreateGenre(req)
}
func (s *GenreService)UpdateGenre(c context.Context, req *pb.GenreCreate) (*pb.Void, error) {
	return s.storage.Genre().UpdateGenre(req)
}
func (s *GenreService)DeleteGenre(c context.Context,id *pb.ById) (*pb.Void, error) {
	return s.storage.Genre().DeleteGenre(id)
}

func (s *GenreService)GetGenre(c context.Context,req *pb.ByName) (*pb.GenreCreate, error) {
	return s.storage.Genre().GetGenre(req)
}

func (s *GenreService)GetAllGenres(c context.Context, req *pb.NameFilter) (*pb.Genres, error) {
	return s.storage.Genre().GetAllGenres(req)
}
func(s *GenreService)GetBooksByGenre(c context.Context, req *pb.GenreId) (*pb.GenreBooks, error) {
	return s.storage.Genre().GetBooksByGenre(req)
}
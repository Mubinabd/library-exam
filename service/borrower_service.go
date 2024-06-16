package service

import (
	"context"

	pb "github.com/Mubinabd/library-exam/genproto"
	"github.com/Mubinabd/library-exam/storage"
	"github.com/google/uuid"
)

type BorrowerService struct {
	storage storage.StorageI
	pb.UnimplementedBorrowerServiceServer
}

func NewBorrowerStorage(storage storage.StorageI) *BorrowerService {
	return &BorrowerService{
		storage: storage,
	}
}

func (s *BorrowerService)CreateBorrower(c context.Context, req *pb.BorrowerCreate) (*pb.Borrower, error) {
	id := uuid.NewString()
	req.Id = id
	return s.storage.Borrower().CreateBorrower(req)
}
func (s *BorrowerService)UpdateBorrower(c context.Context, req *pb.BorrowerCreate) (*pb.Void, error) {
	return s.storage.Borrower().UpdateBorrower(req)
}
func (s *BorrowerService)DeleteBorrower(c context.Context,id *pb.ById) (*pb.Void, error) {
	return s.storage.Borrower().DeleteBorrower(id)
}

func (s *BorrowerService)GetBorrower(c context.Context,id *pb.ById) (*pb.Borrower, error) {
	return s.storage.Borrower().GetBorrower(id)
}

func (s *BorrowerService)GetAllBorrowers(c context.Context, req *pb.Void) (*pb.Borrowers, error) {
	return s.storage.Borrower().GetAllBorrowers(req)
}

func(s *BorrowerService)BorrowerBooks(c context.Context, id *pb.UserId) (*pb.BorrowedBooks, error) {	
	return s.storage.Borrower().BorrowerBooks(id)
}

func(s *BorrowerService)GetOverdueBooks(c context.Context, req *pb.OverdueRequest) (*pb.BorrowedBooks, error) {
	return s.storage.Borrower().GetOverdueBooks(req)
}

func(s *BorrowerService)HistoryUser(c context.Context, id *pb.UserId) (*pb.BorrowingHistory, error) {
	return s.storage.Borrower().HistoryUser(id)
}
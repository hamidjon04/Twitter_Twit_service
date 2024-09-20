package twitService

import (
	"context"
	"log/slog"
	pb "twit-service/generated/twit"
	"twit-service/models"
	"twit-service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedTwitServiceClientServer
	Storage storage.Twit
	Logget  *slog.Logger
}

func NewService(logger *slog.Logger, storage storage.Twit) *Service {
	return &Service{
		Storage: storage,
		Logget:  logger,
	}
}

func (s *Service) CreateTwit(ctx context.Context, in *pb.CreateTwitReq) (*pb.CreateTwitResp, error) {
	twit := &models.Twit{
		UserID:  in.UserId,
		Content: in.Content,
		Media:   in.Media,
	}

	resp, err := s.Storage.Twit().CreateTwit(ctx, twit)
	if err != nil {
		s.Logget.Error("error creating twit", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return resp, nil
}

func (s *Service) UpdateTwit(ctx context.Context, in *pb.UpadateReq) (*pb.UpdateTwitResp, error) {
	twit := &models.Twit{
		ID:      in.Id,
		Content: in.Content,
		Media:   in.Media,
		UserID:  in.UserId,
	}

	resp, err := s.Storage.Twit().UpdateTwit(ctx, twit)
	if err != nil {
		s.Logget.Error("error updating twit", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return resp, nil
}

func (s *Service) DeleteTwit(ctx context.Context, in *pb.Id) (*pb.Message, error) {
	err := s.Storage.Twit().DeleteTwit(ctx, in.Id)
	if err != nil {
		s.Logget.Error("error deleting twit", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pb.Message{Message: "twit deleted"}, nil
}

func (s *Service) GetTwits(ctx context.Context, in *pb.GetTwitsReq) (*pb.GetTwitsResp, error) {
	twits, err := s.Storage.Twit().GetTwit(ctx, in.UserId)
	if err != nil {
		s.Logget.Error("error getting twits", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return twits, nil
}

func (s *Service) GetFollowerTwit(ctx context.Context, in *pb.GetTwitsReq) (*pb.GetTwitsResp, error) {
	// twits, err := s.Storage.Twit().GetFollowerTwit(ctx, in.UserId)
	// if err != nil {
	// 	s.Logget.Error("error getting follower twits", err)
	// 	return nil, status.Errorf(codes.Internal, "internal error")
	// }

	return nil, nil
}

func (s *Service) AddLike(ctx context.Context, in *pb.AddLikeReq) (*pb.AddLikeResp, error) {

	resp, err := s.Storage.Like().AddLike(s.Storage.MongoDatabase(), in)
	if err != nil {
		s.Logget.Error("error adding like", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return resp, nil
}

func (s *Service) RemoveLike(ctx context.Context, in *pb.DeleteLikeReq) (*pb.Message, error) {
	resp, err := s.Storage.Like().DeleteLike(s.Storage.MongoDatabase(), in)
	if err != nil {
		s.Logget.Error("error removing like", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return resp, nil
}

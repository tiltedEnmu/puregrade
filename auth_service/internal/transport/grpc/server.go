package grpc

import (
	"context"
	"time"

	"github.com/puregrade/puregrade-auth/internal/entities"
	"github.com/puregrade/puregrade-auth/internal/service"
	pb "github.com/puregrade/puregrade-auth/internal/transport/grpc/proto"
)

type GRPCServer struct {
	services service.Service
}

func NewGRPCServer(services service.Service) *GRPCServer {
	return &GRPCServer{services: services}
}

func (s *GRPCServer) SingIn(ctx context.Context, req *pb.SingInRequest) (*pb.SingInResponse, error) {
	access, refresh, err := s.services.GenerateTokens(121) // call to UserService

	return &pb.SingInResponse{
		Access:  access,
		Refresh: refresh,
	}, err
}

func (s *GRPCServer) SingUp(ctx context.Context, req *pb.SingUpRequest) (*pb.SingUpResponse, error) {
	raw := req.GetUser()
	var user entities.User = entities.User{
		Id:        raw.Id,
		Username:  raw.Username,
		Email:     raw.Email,
		Password:  raw.Password,
		Avatar:    raw.Avatar,
		Banned:    raw.Banned,
		BanReason: raw.BanReason,
		Status:    raw.Status,
		Followers: raw.Followers,
		Roles:     raw.Roles,
		CreatedAt: time.Unix(raw.CreatedAt, 0),
	}

	id, err := s.services.CreateUser(user)
	if err != nil {
		var res *pb.SingUpResponse
		return res, err
	}

	access, refresh, err := s.services.GenerateTokens(id)

	return &pb.SingUpResponse{
		Access:  access,
		Refresh: refresh,
	}, err
}

func (s *GRPCServer) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	id, err := s.services.GetUserId(req.GetRefresh())
	if err != nil {
		return &pb.RefreshResponse{
			Access:  "",
			Refresh: "",
		}, err
	}

	access, refresh, err := s.services.GenerateTokens(id)
	if err != nil {
		return &pb.RefreshResponse{
			Access:  "",
			Refresh: "",
		}, err
	}

	return &pb.RefreshResponse{
		Access:  access,
		Refresh: refresh,
	}, nil
}

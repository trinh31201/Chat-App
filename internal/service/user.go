package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	userV1 "github.com/yourusername/chat-app/api/user/v1"
	"github.com/yourusername/chat-app/internal/biz"
)

// UserService implements the user service
type UserService struct {
	userV1.UnimplementedUserServiceServer

	uc  *biz.UserUseCase
	log *log.Helper
}

// NewUserService creates a new user service
func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/user")),
	}
}

// Register creates a new user account
func (s *UserService) Register(ctx context.Context, req *userV1.RegisterRequest) (*userV1.RegisterResponse, error) {
	user, token, err := s.uc.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return &userV1.RegisterResponse{
		User: &userV1.User{
			Id:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			AvatarUrl: user.AvatarURL,
			Status:    user.Status,
			LastSeen:  user.LastSeen.Unix(),
			CreatedAt: user.CreatedAt.Unix(),
		},
		Token: token,
	}, nil
}

// Login authenticates a user and returns a token
func (s *UserService) Login(ctx context.Context, req *userV1.LoginRequest) (*userV1.LoginResponse, error) {
	user, token, err := s.uc.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return &userV1.LoginResponse{
		User: &userV1.User{
			Id:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			AvatarUrl: user.AvatarURL,
			Status:    user.Status,
			LastSeen:  user.LastSeen.Unix(),
			CreatedAt: user.CreatedAt.Unix(),
		},
		Token: token,
	}, nil
}

// GetUser retrieves user information by ID
func (s *UserService) GetUser(ctx context.Context, req *userV1.GetUserRequest) (*userV1.User, error) {
	user, err := s.uc.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &userV1.User{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		AvatarUrl: user.AvatarURL,
		Status:    user.Status,
		LastSeen:  user.LastSeen.Unix(),
		CreatedAt: user.CreatedAt.Unix(),
	}, nil
}

// UpdateStatus updates the user's online status
func (s *UserService) UpdateStatus(ctx context.Context, req *userV1.UpdateStatusRequest) (*userV1.UpdateStatusResponse, error) {
	err := s.uc.UpdateStatus(ctx, req.UserId, req.Status)
	if err != nil {
		return nil, err
	}

	return &userV1.UpdateStatusResponse{
		Success: true,
		Status:  req.Status,
	}, nil
}
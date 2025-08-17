package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	chatV1 "github.com/yourusername/chat-app/api/chat/v1"
	"github.com/yourusername/chat-app/internal/biz"
)

// ChatService implements the chat/messaging service
type ChatService struct {
	chatV1.UnimplementedChatServiceServer

	uc  *biz.ChatUseCase
	log *log.Helper
}

// NewChatService creates a new chat service
func NewChatService(uc *biz.ChatUseCase, logger log.Logger) *ChatService {
	return &ChatService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/chat")),
	}
}

// SendMessage sends a message to a room
func (s *ChatService) SendMessage(ctx context.Context, req *chatV1.SendMessageRequest) (*chatV1.Message, error) {
	// Get user ID from context (set by authentication middleware)
	userID, err := s.getUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	message, err := s.uc.SendMessage(ctx, userID, req)
	if err != nil {
		return nil, err
	}

	return &chatV1.Message{
		Id:        message.ID,
		RoomId:    message.RoomID,
		UserId:    message.UserID,
		Username:  message.Username,
		Content:   message.Content,
		Type:      message.Type,
		IsEdited:  message.IsEdited,
		CreatedAt: message.CreatedAt.Unix(),
	}, nil
}

// TODO: Additional chat service methods will be implemented as needed
// GetMessage, ListMessages, EditMessage, DeleteMessage, MarkMessageAsRead, GetUnreadMessages
// These require corresponding protocol buffer message definitions and data layer implementations

// getUserIDFromContext extracts user ID from request context
// This would be set by an authentication middleware
func (s *ChatService) getUserIDFromContext(ctx context.Context) (int64, error) {
	userID, ok := ctx.Value("user_id").(int64)
	if !ok {
		return 0, biz.ErrUserNotFound
	}
	
	if userID <= 0 {
		return 0, biz.ErrUserNotFound
	}
	
	return userID, nil
}
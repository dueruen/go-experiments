package writting

import (
	"context"

	listing "github.com/dueruen/go-experiments/go-kit/chat/pkg/listing"
)

type Service interface {
	WriteToChat(ctx context.Context, chatItem ChatItem) (listing.ChatItem, error)
}

type Repository interface {
	WriteToChat(chatItem ChatItem) (listing.ChatItem, error)
}

type service struct {
	chat_repo Repository
}

func NewService(chat_repo Repository) Service {
	return &service{chat_repo}
}

func (service *service) WriteToChat(ctx context.Context, chatItem ChatItem) (listing.ChatItem, error) {
	return service.chat_repo.WriteToChat(chatItem)
}

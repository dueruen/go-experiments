package transport

import (
	"github.com/dueruen/go-experiments/go-kit/chat/pkg/listing"
	"github.com/dueruen/go-experiments/go-kit/chat/pkg/writting"
)

type (
	WriteToChatRequest struct {
		ChatItem writting.ChatItem
	}

	WriteToChatResponse struct {
		ChatItem listing.ChatItem
	}

	// GetChatRequest struct{}

	// GetChatResponse struct {
	// 	Chat []chat.ChatItem
	// }
)

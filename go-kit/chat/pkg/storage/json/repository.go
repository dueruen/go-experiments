package json

import (
	"path"
	"runtime"

	"github.com/dueruen/go-experiments/go-kit/chat/pkg/listing"
	"github.com/dueruen/go-experiments/go-kit/chat/pkg/writting"
	uuid "github.com/gofrs/uuid"
	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	dir            = "/data/"
	CollectionChat = "chat"
)

type Storage struct {
	db *scribble.Driver
}

func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)
	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (storage *Storage) WriteToChat(chatItem writting.ChatItem) (listing.ChatItem, error) {
	uuid, _ := uuid.NewV4()
	newMessage := ChatItem{
		ID:      uuid.String(),
		Auther:  chatItem.Auther,
		Message: chatItem.Message,
	}

	createdItem := new(listing.ChatItem)
	if err := storage.db.Write(CollectionChat, newMessage.ID, newMessage); err != nil {
		return *createdItem, err
	}
	createdItem.ID = newMessage.ID
	createdItem.Auther = newMessage.Auther
	createdItem.Message = newMessage.Message
	return *createdItem, nil
}

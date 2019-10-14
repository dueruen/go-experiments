package json

import (
	"encoding/json"
	"path"
	"runtime"

	om "github.com/dueruen/go-experiments/go-graphql/service/user/pkg"
	uuid "github.com/gofrs/uuid"
	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	dir             = "/data/"
	CollectionUsers = "users"
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

func (storage *Storage) CreateNewUser(user om.CreateUserRequest) (om.User, error) {
	uuid, _ := uuid.NewV4()
	newUser := om.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		ID:        uuid.String(),
	}

	if err := storage.db.Write(CollectionUsers, newUser.ID, newUser); err != nil {
		return om.User{}, err
	}
	return newUser, nil
}

func (storage *Storage) ListAllUsers() ([]om.User, error) {
	list := []om.User{}

	records, err := storage.db.ReadAll(CollectionUsers)
	if err != nil {
		return list, nil
	}

	for _, r := range records {
		var user om.User

		if err := json.Unmarshal([]byte(r), &user); err != nil {
			// err handling omitted for simplicity
			return list, nil
		}

		list = append(list, user)
	}
	return list, nil
}

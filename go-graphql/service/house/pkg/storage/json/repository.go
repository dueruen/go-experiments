package json

import (
	"encoding/json"
	"path"
	"runtime"

	om "github.com/dueruen/go-experiments/go-graphql/service/house/pkg"
	uuid "github.com/gofrs/uuid"
	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	dir              = "/data/"
	CollectionHouses = "houses"
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

func (storage *Storage) CreateNewHouse(house om.CreateHouseRequest) (om.House, error) {
	uuid, _ := uuid.NewV4()
	newHouse := om.House{
		Address: house.Address,
		OwnerID: house.OwnerID,
		Age:     house.Age,
		ID:      uuid.String(),
	}

	if err := storage.db.Write(CollectionHouses, newHouse.ID, newHouse); err != nil {
		return om.House{}, err
	}
	return newHouse, nil
}

func (storage *Storage) ListAllHouses() ([]om.House, error) {
	list := []om.House{}

	records, err := storage.db.ReadAll(CollectionHouses)
	if err != nil {
		return list, nil
	}

	for _, r := range records {
		var house om.House

		if err := json.Unmarshal([]byte(r), &house); err != nil {
			// err handling omitted for simplicity
			return list, nil
		}

		list = append(list, house)
	}
	return list, nil
}

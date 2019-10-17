package pkg

type HouseService interface {
	CreateHouse(req CreateHouseRequest) (res CreateHouseResponse, err error)
	ListAllHouses(req ListAllHousesRequest) (res ListAllHousesResponse, err error)
}

type Repository interface {
	CreateNewHouse(house CreateHouseRequest) (House, error)
	ListAllHouses() ([]House, error)
}

type House struct {
	Address string
	OwnerID string
	Age     int32
	ID      string
}
type (
	CreateHouseRequest struct {
		Address string
		OwnerID string
		Age     int32
	}

	CreateHouseResponse struct {
		House House
	}

	ListAllHousesRequest struct {
	}

	ListAllHousesResponse struct {
		Houses []House
	}
)

package pkg

type service struct {
	repo Repository
}

func NewService(repo Repository) HouseService {
	return &service{
		repo: repo,
	}
}

func (svc *service) CreateHouse(req CreateHouseRequest) (res CreateHouseResponse, err error) {
	house, err := svc.repo.CreateNewHouse(req)
	return CreateHouseResponse{
		House: house,
	}, err
}

func (svc *service) ListAllHouses(req ListAllHousesRequest) (res ListAllHousesResponse, err error) {
	list, err := svc.repo.ListAllHouses()
	return ListAllHousesResponse{
		Houses: list,
	}, err
}

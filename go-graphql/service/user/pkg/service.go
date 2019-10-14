package user

type service struct {
	repo Repository
}

func NewService(repo Repository) UserService {
	return &service{
		repo: repo,
	}
}

func (svc *service) CreateUser(req CreateUserRequest) (res CreateUserResponse, err error) {
	user, err := svc.repo.CreateNewUser(req)
	return CreateUserResponse{
		User: user,
	}, err
}

func (svc *service) ListAllUsers(req ListAllUsersRequest) (res ListAllUsersResponse, err error) {
	list, err := svc.repo.ListAllUsers()
	return ListAllUsersResponse{
		Users: list,
	}, err
}

package user

type UserService interface {
	CreateUser(req CreateUserRequest) (res CreateUserResponse, err error)
	ListAllUsers(req ListAllUsersRequest) (res ListAllUsersResponse, err error)
}

type Repository interface {
	CreateNewUser(user CreateUserRequest) (User, error)
	ListAllUsers() ([]User, error)
}

type User struct {
	FirstName string
	LastName  string
	Age       int32
	ID        string
}
type (
	CreateUserRequest struct {
		FirstName string
		LastName  string
		Age       int32
	}

	CreateUserResponse struct {
		User User
	}

	ListAllUsersRequest struct {
	}

	ListAllUsersResponse struct {
		Users []User
	}
)

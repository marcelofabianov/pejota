package port

// Repository

type UserRepository interface {
	GetUsersRepository
	GetUserRepository
	CreateUserRepository
	DeleteUserRepository
	UpdateUserRepository
	DisableUserLoginRepository
}

// Service

type UserServiceApplication interface {
	GetUsersServiceApplication
	GetUserServiceApplication
	CreateUserServiceApplication
	DeleteUserServiceApplication
	UpdateUserServiceApplication
	DisableUserLoginServiceApplication
}

package port

// Repository

type UserRepository interface {
	GetUsersRepository
	GetUserRepository
	CreateUserRepository
	DeleteUserRepository
	UpdateUserRepository
	DisableUserLoginRepository
	EnableUserLoginRepository
}

// Service

type UserServiceApplication interface {
	GetUsersServiceApplication
	GetUserServiceApplication
	CreateUserServiceApplication
	DeleteUserServiceApplication
	UpdateUserServiceApplication
	DisableUserLoginServiceApplication
	EnableUserLoginServiceApplication
}

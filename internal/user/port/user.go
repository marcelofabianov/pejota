package port

// Repository

type UserRepository interface {
	GetUserRepository
	CreateUserRepository
	DeleteUserRepository
	UpdateUserRepository
}

// Service

type UserServiceApplication interface {
	GetUserServiceApplication
	CreateUserServiceApplication
	DeleteUserServiceApplication
	UpdateUserServiceApplication
}

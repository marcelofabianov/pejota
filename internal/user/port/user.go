package port

// Repository

type UserRepository interface {
	GetUserRepository
	CreateUserRepository
	DeleteUserRepository
}

// Service

type UserServiceApplication interface {
	GetUserServiceApplication
	CreateUserServiceApplication
	DeleteUserServiceApplication
}

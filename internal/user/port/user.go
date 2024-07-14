package port

// Repository

type UserRepository interface {
	GetUserRepository
	CreateUserRepository
}

// Service

type UserServiceApplication interface {
	GetUserServiceApplication
	CreateUserServiceApplication
}

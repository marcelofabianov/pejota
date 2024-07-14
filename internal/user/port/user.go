package port

// Repository

type UserRepository interface {
	GetUserRepository
}

// Service

type UserServiceApplication interface {
	GetUserServiceApplication
}

// PKG

type PasswordHasher interface {
	Hash(data string) (string, error)
	Compare(data, encodedHash string) (bool, error)
}

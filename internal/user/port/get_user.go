package port

// Service Application

type GetUserInputServiceApplication struct {
	PublicID string `json:"public_id" validate:"required,uuid"`
}

// using GetUsersOutputServiceApplication from get_users.go
type GetUserOutputServiceApplication struct {
	PublicID     string `json:"public_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	LoginEnabled bool   `json:"login_enabled"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type GetUserServiceApplication interface {
	GetUser(input GetUserInputServiceApplication) (GetUserOutputServiceApplication, error)
}

// UseCase

type GetUserInputUseCase struct {
	GetUserInputServiceApplication
}

type GetUserOutputUseCase struct {
	GetUserOutputRepository
}

type GetUserUseCase interface {
	Execute(input GetUserInputUseCase) (GetUserOutputUseCase, error)
}

// Repository

type GetUserInputRepository struct {
	PublicID string
}

type GetUserOutputRepository struct {
	PublicID     string
	Name         string
	Email        string
	Role         string
	LoginEnabled bool
	CreatedAt    string
	UpdatedAt    string
}

type GetUserRepository interface {
	GetUser(input GetUserInputRepository) (GetUserOutputRepository, error)
}

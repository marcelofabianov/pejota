package port

// Service

type GetUserInputServiceApplication struct {
	PublicID string `json:"public_id" validate:"required,uuid"`
}

type GetUserOutputServiceApplication struct {
	PublicID     string `json:"public_id" validate:"required,uuid"`
	Name         string `json:"name" validate:"required,string,min=3,max=255"`
	Email        string `json:"email" validate:"required,email"`
	Role         string `json:"role" validate:"required,role"`
	LoginEnabled bool   `json:"login_enabled" validate:"required,bool"`
	CreatedAt    string `json:"created_at" validate:"required,datetime"`
	UpdatedAt    string `json:"updated_at" validate:"required,datetime"`
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

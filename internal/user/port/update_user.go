package port

type UpdateUserInputServiceApplication struct {
	PublicID string `json:"public_id" validate:"required,uuid"`
	Name     string `json:"name" validate:"required,string,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required,role"`
}

type UpdateUserOutputServiceApplication struct {
	PublicID  string `json:"public_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateUserServiceApplication interface {
	UpdateUser(input UpdateUserInputServiceApplication) (UpdateUserOutputServiceApplication, error)
}

// UseCase

type UpdateUserInputUseCase struct {
	UpdateUserInputServiceApplication
}

type UpdateUserOutputUseCase struct {
	UpdateUserOutputRepository
}

type UpdateUserUseCase interface {
	Execute(input UpdateUserInputUseCase) (UpdateUserOutputUseCase, error)
}

// Repository

type UpdateUserInputRepository struct {
	ID        int64
	Name      string `validate:"required,string,min=3,max=255"`
	Email     string `validate:"required,email"`
	Role      string `validate:"required,role"`
	UpdatedAt string `validate:"required,datetime"`
}

type UpdateUserOutputRepository struct {
	PublicID  string
	Name      string
	Email     string
	Role      string
	CreatedAt string
	UpdatedAt string
}

type UpdateUserRepository interface {
	FindUserIDByPublicID(input FindUserIDByPublicIDInputRepository) (FindUserIDByPublicIDOutputRepository, error)
	UpdateUser(input UpdateUserInputRepository) (UpdateUserOutputRepository, error)
}

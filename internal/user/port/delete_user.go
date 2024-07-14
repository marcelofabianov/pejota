package port

type DeleteUserInputServiceApplication struct {
	PublicID string `json:"public_id" validate:"required,uuid"`
}

type DeleteUserOutputServiceApplication struct {
	Success bool `json:"success"`
}

type DeleteUserServiceApplication interface {
	DeleteUser(input DeleteUserInputServiceApplication) (DeleteUserOutputServiceApplication, error)
}

// UseCase

type DeleteUserInputUseCase struct {
	DeleteUserInputServiceApplication
}

type DeleteUserOutputUseCase struct {
	DeleteUserOutputRepository
}

type DeleteUserUseCase interface {
	Execute(input DeleteUserInputUseCase) (DeleteUserOutputUseCase, error)
}

// Repository

type DeleteUserInputRepository struct {
	PublicID string
}

type DeleteUserOutputRepository struct {
	Success bool
}

type FindUserInputRepository struct {
	PublicID string
}

type FindUserOutputRepository struct {
	ID string
}

type DeleteUserRepository interface {
	GetUserById(input FindUserInputRepository) (FindUserOutputRepository, error)
	DeleteUser(input DeleteUserInputRepository) (DeleteUserOutputRepository, error)
}

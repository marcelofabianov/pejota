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

type FindUserIDByPublicIDInputRepository struct {
	PublicID string
}

type FindUserIDByPublicIDOutputRepository struct {
	ID int64
}

type DeleteUserInputRepository struct {
	FindUserIDByPublicIDOutputRepository
	DeletedAt string
}

type DeleteUserOutputRepository struct {
	Success bool
}

type DeleteUserRepository interface {
	FindUserIDByPublicID(input FindUserIDByPublicIDInputRepository) (FindUserIDByPublicIDOutputRepository, error)
	DeleteUser(input DeleteUserInputRepository) (DeleteUserOutputRepository, error)
}

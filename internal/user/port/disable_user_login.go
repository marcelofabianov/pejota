package port

// Service Application

type DisableUserLoginInputServiceApplication struct {
	PublicID string `json:"public_id" validate:"required,uuid"`
}

type DisableUserLoginOutputServiceApplication struct {
	PublicID     string `json:"public_id"`
	LoginEnabled bool   `json:"login_enabled"`
}

type DisableUserLoginServiceApplication interface {
	DisableUserLogin(input DisableUserLoginInputServiceApplication) (DisableUserLoginOutputServiceApplication, error)
}

// UseCase

type DisableUserLoginInputUseCase struct {
	DisableUserLoginInputServiceApplication
}

type DisableUserLoginOutputUseCase struct {
	DisableUserLoginOutputRepository
}

type DisableUserLoginUseCase interface {
	Execute(input DisableUserLoginInputUseCase) (DisableUserLoginOutputUseCase, error)
}

// Repository

type DisableUserLoginInputRepository struct {
	PublicID  string
	UpdatedAt string `validate:"required,datetime"`
}

type DisableUserLoginOutputRepository struct {
	PublicID     string
	LoginEnabled bool
}

type DisableUserLoginRepository interface {
	DisableUserLogin(input DisableUserLoginInputRepository) (DisableUserLoginOutputRepository, error)
}

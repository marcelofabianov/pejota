package port

type EnableUserLoginInputServiceApplication struct {
	PublicID string `json:"public_id" validate:"required,uuid"`
}

type EnableUserLoginOutputServiceApplication struct {
	PublicID     string `json:"public_id"`
	LoginEnabled bool   `json:"login_enabled"`
}

type EnableUserLoginServiceApplication interface {
	EnableUserLogin(input EnableUserLoginInputServiceApplication) (EnableUserLoginOutputServiceApplication, error)
}

// UseCase

type EnableUserLoginInputUseCase struct {
	EnableUserLoginInputServiceApplication
}

type EnableUserLoginOutputUseCase struct {
	EnableUserLoginOutputRepository
}

type EnableUserLoginUseCase interface {
	Execute(input EnableUserLoginInputUseCase) (EnableUserLoginOutputUseCase, error)
}

// Repository

type EnableUserLoginInputRepository struct {
	PublicID  string
	UpdatedAt string `validate:"required,datetime"`
}

type EnableUserLoginOutputRepository struct {
	PublicID     string
	LoginEnabled bool
}

type EnableUserLoginRepository interface {
	EnableUserLogin(input EnableUserLoginInputRepository) (EnableUserLoginOutputRepository, error)
}

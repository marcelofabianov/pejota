package port

type GetUserInputService struct {
	PublicID string `json:"public_id" validate:"required,uuid"`
}

type GetUserOutputService struct {
	PublicID     string `json:"public_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	LoginEnabled bool   `json:"login_enabled"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type GetUserInputUseCase struct {
	GetUserInputService
}

type GetUserOutputUseCase struct {
	GetUserOutputService
}

type GetUserUseCase interface {
	Execute(input GetUserInputUseCase) (GetUserOutputUseCase, error)
}

type GetUserService interface {
	GetUser(input GetUserInputService) (GetUserOutputService, error)
}

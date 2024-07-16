package port

type GetUsersInputServiceApplication struct {
	Page   int     `json:"page" validate:"required,number"`
	Limit  int     `json:"limit" validate:"required,number"`
	Order  string  `json:"order" validate:"required,string"`
	Search *string `json:"search"`
}

type GetUsersOutputServiceApplication struct {
	Users  []GetUserOutputServiceApplication `json:"users"`
	Totals int                               `json:"totals"`
}

type GetUsersServiceApplication interface {
	GetUsers(input GetUsersInputServiceApplication) (GetUsersOutputServiceApplication, error)
}

// UseCase

type GetUsersInputUseCase struct {
	GetUsersInputServiceApplication
}

type GetUsersOutputUseCase struct {
	GetUsersOutputServiceApplication
}

type GetUsersUseCase interface {
	Execute(input GetUsersInputUseCase) (GetUsersOutputUseCase, error)
}

// Repository

type GetUsersInputRepository struct {
	Page   int
	Limit  int
	Order  string
	Search string
}

type GetUsersOutputRepository struct {
	Users  []GetUserOutputRepository
	Totals int
}

type GetUsersRepository interface {
	GetUsers(input GetUsersInputRepository) (GetUsersOutputRepository, error)
}

package port

// PKG

type PasswordHasher interface {
	Hash(data string) (string, error)
	Compare(data, encodedHash string) (bool, error)
}

// Service Application

type CreateUserInputServiceApplication struct {
	Name     string `json:"name" validate:"required,string,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,string,min=6,max=255"`
	Role     string `json:"role" validate:"required,role"`
}

type CreateUserOutputServiceApplication struct {
	PublicID  string `json:"public_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
}

type CreateUserServiceApplication interface {
	CreateUser(input CreateUserInputServiceApplication) (CreateUserOutputServiceApplication, error)
}

// UseCase

type CreateUserInputUseCase struct {
	CreateUserInputServiceApplication
}

type CreateUserOutputUseCase struct {
	CreateUserOutputRepository
}

type CreateUserUseCase interface {
	Execute(input CreateUserInputUseCase) (CreateUserOutputUseCase, error)
}

// Repository

type CreateUserInputRepository struct {
	Name     string `validate:"required,string,min=3,max=255"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,string,min=6,max=255"`
	Role     string `validate:"required,role"`
}

type CreateUserOutputRepository struct {
	PublicID  string
	Name      string
	Email     string
	Role      string
	CreatedAt string
	UpdateAt  string
}

type CreateUserRepository interface {
	CreateUser(input CreateUserInputRepository) (CreateUserOutputRepository, error)
}

package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/marcelofabianov/pejota/internal/user/port"
)

type CreateUserUseCase struct {
	hasher     port.PasswordHasher
	repository port.CreateUserRepository
}

func NewCreateUserUseCase(hasher port.PasswordHasher, repository port.CreateUserRepository) port.CreateUserUseCase {
	return &CreateUserUseCase{hasher, repository}
}

func (u *CreateUserUseCase) Execute(input port.CreateUserInputUseCase) (port.CreateUserOutputUseCase, error) {
	uuid := uuid.New().String()
	hashedPassword, err := u.hasher.Hash(input.Password)
	if err != nil {
		return port.CreateUserOutputUseCase{}, err
	}

	repositoryInput := port.CreateUserInputRepository{
		PublicID:     uuid,
		Name:         input.Name,
		Email:        input.Email,
		Password:     hashedPassword,
		LoginEnabled: true,
		Role:         input.Role,
		CreatedAt:    time.Now().Format(time.RFC3339),
		UpdatedAt:    time.Now().Format(time.RFC3339),
	}

	repositoryOutput, err := u.repository.CreateUser(repositoryInput)
	if err != nil {
		return port.CreateUserOutputUseCase{}, err
	}

	output := port.CreateUserOutputUseCase{
		CreateUserOutputRepository: port.CreateUserOutputRepository{
			PublicID:     repositoryOutput.PublicID,
			Name:         repositoryOutput.Name,
			Email:        repositoryOutput.Email,
			LoginEnabled: repositoryOutput.LoginEnabled,
			Role:         repositoryOutput.Role,
			CreatedAt:    repositoryOutput.CreatedAt,
			UpdatedAt:    repositoryOutput.UpdatedAt,
		},
	}

	return output, nil
}

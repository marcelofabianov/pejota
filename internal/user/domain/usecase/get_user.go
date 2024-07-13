package usecase

import "github.com/marcelofabianov/pejota/internal/user/port"

type GetUserUseCase struct {
	repository port.GetUserRepository
}

func NewGetUserUseCase(repository port.GetUserRepository) port.GetUserUseCase {
	return &GetUserUseCase{repository}
}

func (u *GetUserUseCase) Execute(input port.GetUserInputUseCase) (port.GetUserOutputUseCase, error) {
	repositoryInput := port.GetUserInputRepository{
		PublicID: input.PublicID,
	}

	repositoryOutput, err := u.repository.GetUser(repositoryInput)
	if err != nil {
		return port.GetUserOutputUseCase{}, err
	}

	output := port.GetUserOutputUseCase{
		GetUserOutputRepository: port.GetUserOutputRepository{
			PublicID:     repositoryOutput.PublicID,
			Name:         repositoryOutput.Name,
			Email:        repositoryOutput.Email,
			Role:         repositoryOutput.Role,
			LoginEnabled: repositoryOutput.LoginEnabled,
			CreatedAt:    repositoryOutput.CreatedAt,
			UpdatedAt:    repositoryOutput.UpdatedAt,
		},
	}

	return output, nil
}

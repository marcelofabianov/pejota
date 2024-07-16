package usecase

import (
	"time"

	"github.com/marcelofabianov/pejota/internal/user/port"
)

type EnableUserLoginUseCase struct {
	repository port.EnableUserLoginRepository
}

func NewEnableUserLoginUseCase(repository port.EnableUserLoginRepository) port.EnableUserLoginUseCase {
	return &EnableUserLoginUseCase{repository: repository}
}

func (u *EnableUserLoginUseCase) Execute(input port.EnableUserLoginInputUseCase) (port.EnableUserLoginOutputUseCase, error) {
	repoInput := port.EnableUserLoginInputRepository{
		PublicID:  input.PublicID,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	repoOutput, err := u.repository.EnableUserLogin(repoInput)
	if err != nil {
		return port.EnableUserLoginOutputUseCase{}, err
	}

	output := port.EnableUserLoginOutputUseCase{
		EnableUserLoginOutputRepository: port.EnableUserLoginOutputRepository{
			PublicID:     repoOutput.PublicID,
			LoginEnabled: repoOutput.LoginEnabled,
		},
	}

	return output, nil
}

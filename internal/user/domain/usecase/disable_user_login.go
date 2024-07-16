package usecase

import (
	"time"

	"github.com/marcelofabianov/pejota/internal/user/port"
)

type DisableUserLoginUseCase struct {
	repository port.DisableUserLoginRepository
}

func NewDisableUserLoginUseCase(repository port.DisableUserLoginRepository) port.DisableUserLoginUseCase {
	return &DisableUserLoginUseCase{repository: repository}
}

func (u *DisableUserLoginUseCase) Execute(input port.DisableUserLoginInputUseCase) (port.DisableUserLoginOutputUseCase, error) {
	repoInput := port.DisableUserLoginInputRepository{
		PublicID:  input.PublicID,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	repoOutput, err := u.repository.DisableUserLogin(repoInput)
	if err != nil {
		return port.DisableUserLoginOutputUseCase{}, err
	}

	output := port.DisableUserLoginOutputUseCase{
		DisableUserLoginOutputRepository: port.DisableUserLoginOutputRepository{
			PublicID:     repoOutput.PublicID,
			LoginEnabled: repoOutput.LoginEnabled,
		},
	}

	return output, nil
}

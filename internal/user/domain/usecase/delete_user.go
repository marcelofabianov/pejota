package usecase

import (
	"time"

	"github.com/marcelofabianov/pejota/internal/user/port"
)

type DeleteUserUseCase struct {
	repository port.DeleteUserRepository
}

func NewDeleteUserUseCase(repository port.DeleteUserRepository) port.DeleteUserUseCase {
	return &DeleteUserUseCase{
		repository: repository,
	}
}

func (uc *DeleteUserUseCase) Execute(input port.DeleteUserInputUseCase) (port.DeleteUserOutputUseCase, error) {
	// find user
	findUserInput := port.FindUserIDByPublicIDInputRepository{
		PublicID: input.PublicID,
	}

	findUserOutput, err := uc.repository.FindUserIDByPublicID(findUserInput)
	if err != nil {
		return port.DeleteUserOutputUseCase{}, err
	}

	// delete user
	deleteUserInput := port.DeleteUserInputRepository{
		FindUserIDByPublicIDOutputRepository: findUserOutput,
		DeletedAt:                            time.Now().Format(time.RFC3339),
	}

	deleteUserOutput, err := uc.repository.DeleteUser(deleteUserInput)
	if err != nil {
		return port.DeleteUserOutputUseCase{}, err
	}

	output := port.DeleteUserOutputUseCase{
		DeleteUserOutputRepository: port.DeleteUserOutputRepository{
			Success: deleteUserOutput.Success,
		},
	}

	return output, nil
}

package usecase

import "github.com/marcelofabianov/pejota/internal/user/port"

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
	findUserInput := port.FindUserInputRepository{
		PublicID: input.PublicID,
	}

	findUserOutput, err := uc.repository.FindUser(findUserInput)
	if err != nil {
		return port.DeleteUserOutputUseCase{}, err
	}

	// delete user
	deleteUserInput := port.DeleteUserInputRepository{
		FindUserOutputRepository: findUserOutput,
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

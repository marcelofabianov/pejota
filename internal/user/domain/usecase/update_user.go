package usecase

import (
	"time"

	"github.com/marcelofabianov/pejota/internal/user/port"
)

type UpdateUserUseCase struct {
	repository port.UpdateUserRepository
}

func NewUpdateUserUseCase(repository port.UpdateUserRepository) port.UpdateUserUseCase {
	return &UpdateUserUseCase{repository: repository}
}

func (u *UpdateUserUseCase) Execute(input port.UpdateUserInputUseCase) (port.UpdateUserOutputUseCase, error) {
	// Find user by public id
	findUserIDByPublicIDInputRepository := port.FindUserIDByPublicIDInputRepository{
		PublicID: input.PublicID,
	}
	findUserIDByPublicIDOutputRepository, err := u.repository.FindUserIDByPublicID(findUserIDByPublicIDInputRepository)
	if err != nil {
		return port.UpdateUserOutputUseCase{}, err
	}

	// Update user
	updateUserInputRepository := port.UpdateUserInputRepository{
		ID:        findUserIDByPublicIDOutputRepository.ID,
		Name:      input.Name,
		Email:     input.Email,
		Role:      input.Role,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	updateUserOutputRepository, err := u.repository.UpdateUser(updateUserInputRepository)
	if err != nil {
		return port.UpdateUserOutputUseCase{}, err
	}

	output := port.UpdateUserOutputUseCase{
		UpdateUserOutputRepository: port.UpdateUserOutputRepository{
			PublicID:  updateUserOutputRepository.PublicID,
			Name:      updateUserOutputRepository.Name,
			Email:     updateUserOutputRepository.Email,
			Role:      updateUserOutputRepository.Role,
			CreatedAt: updateUserOutputRepository.CreatedAt,
			UpdatedAt: updateUserOutputRepository.UpdatedAt,
		},
	}

	return output, nil
}

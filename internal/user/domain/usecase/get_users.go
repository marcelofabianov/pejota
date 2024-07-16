package usecase

import "github.com/marcelofabianov/pejota/internal/user/port"

type GetUsersUseCase struct {
	repository port.GetUsersRepository
}

func NewGetUsersUseCase(repository port.GetUsersRepository) port.GetUsersUseCase {
	return &GetUsersUseCase{repository}
}

func (u *GetUsersUseCase) Execute(input port.GetUsersInputUseCase) (port.GetUsersOutputUseCase, error) {
	repositoryInput := port.GetUsersInputRepository{
		Page:   input.Page,
		Limit:  input.Limit,
		Order:  input.Order,
		Search: input.Search,
	}

	repositoryOutput, err := u.repository.GetUsers(repositoryInput)
	if err != nil {
		return port.GetUsersOutputUseCase{}, err
	}

	output := port.GetUsersOutputUseCase{
		GetUsersOutputServiceApplication: port.GetUsersOutputServiceApplication{
			Users: []port.GetUserOutputServiceApplication{},
			Total: repositoryOutput.Total,
		},
	}

	for _, user := range repositoryOutput.Users {
		output.Users = append(output.Users, port.GetUserOutputServiceApplication{
			PublicID:     user.PublicID,
			Name:         user.Name,
			Email:        user.Email,
			Role:         user.Role,
			LoginEnabled: user.LoginEnabled,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		})
	}

	return output, nil
}

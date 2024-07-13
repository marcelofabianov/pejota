package application

import "github.com/marcelofabianov/pejota/internal/user/port"

type UserServiceApplication struct {
	getUser port.GetUserUseCase
}

func NewUserServiceApplication(getUser port.GetUserUseCase) port.UserServiceApplication {
	return &UserServiceApplication{getUser}
}

func (s *UserServiceApplication) GetUser(input port.GetUserInputServiceApplication) (port.GetUserOutputServiceApplication, error) {
	inputUseCase := port.GetUserInputUseCase{GetUserInputServiceApplication: input}

	outputUseCase, err := s.getUser.Execute(inputUseCase)
	if err != nil {
		return port.GetUserOutputServiceApplication{}, err
	}

	return port.GetUserOutputServiceApplication{
		PublicID:     outputUseCase.PublicID,
		Name:         outputUseCase.Name,
		Email:        outputUseCase.Email,
		Role:         outputUseCase.Role,
		LoginEnabled: outputUseCase.LoginEnabled,
		CreatedAt:    outputUseCase.CreatedAt,
		UpdatedAt:    outputUseCase.UpdatedAt,
	}, nil
}

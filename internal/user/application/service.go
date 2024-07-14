package application

import "github.com/marcelofabianov/pejota/internal/user/port"

type UserServiceApplication struct {
	getUser    port.GetUserUseCase
	createUser port.CreateUserUseCase
}

func NewUserServiceApplication(getUser port.GetUserUseCase, createUser port.CreateUserUseCase) port.UserServiceApplication {
	return &UserServiceApplication{getUser, createUser}
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

func (s *UserServiceApplication) CreateUser(input port.CreateUserInputServiceApplication) (port.CreateUserOutputServiceApplication, error) {
	inputUseCase := port.CreateUserInputUseCase{CreateUserInputServiceApplication: input}

	outputUseCase, err := s.createUser.Execute(inputUseCase)
	if err != nil {
		return port.CreateUserOutputServiceApplication{}, err
	}

	return port.CreateUserOutputServiceApplication{
		PublicID:  outputUseCase.PublicID,
		Name:      outputUseCase.Name,
		Email:     outputUseCase.Email,
		Role:      outputUseCase.Role,
		CreatedAt: outputUseCase.CreatedAt,
		UpdatedAt: outputUseCase.UpdatedAt,
	}, nil
}

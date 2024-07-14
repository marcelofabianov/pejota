package application

import "github.com/marcelofabianov/pejota/internal/user/port"

type UserServiceApplication struct {
	getUser    port.GetUserUseCase
	createUser port.CreateUserUseCase
	deleteUser port.DeleteUserUseCase
	updateUser port.UpdateUserUseCase
}

func NewUserServiceApplication(getUser port.GetUserUseCase, createUser port.CreateUserUseCase, deleteUser port.DeleteUserUseCase, updateUser port.UpdateUserUseCase) port.UserServiceApplication {
	return &UserServiceApplication{getUser, createUser, deleteUser, updateUser}
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

func (s *UserServiceApplication) DeleteUser(input port.DeleteUserInputServiceApplication) (port.DeleteUserOutputServiceApplication, error) {
	inputUseCase := port.DeleteUserInputUseCase{DeleteUserInputServiceApplication: input}

	outputUseCase, err := s.deleteUser.Execute(inputUseCase)
	if err != nil {
		return port.DeleteUserOutputServiceApplication{}, err
	}

	return port.DeleteUserOutputServiceApplication{
		Success: outputUseCase.Success,
	}, nil
}

func (s *UserServiceApplication) UpdateUser(input port.UpdateUserInputServiceApplication) (port.UpdateUserOutputServiceApplication, error) {
	inputUseCase := port.UpdateUserInputUseCase{UpdateUserInputServiceApplication: input}

	outputUseCase, err := s.updateUser.Execute(inputUseCase)
	if err != nil {
		return port.UpdateUserOutputServiceApplication{}, err
	}

	return port.UpdateUserOutputServiceApplication{
		PublicID:  outputUseCase.PublicID,
		Name:      outputUseCase.Name,
		Email:     outputUseCase.Email,
		Role:      outputUseCase.Role,
		CreatedAt: outputUseCase.CreatedAt,
		UpdatedAt: outputUseCase.UpdatedAt,
	}, nil
}

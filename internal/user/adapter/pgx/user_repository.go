package pgx

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/marcelofabianov/pejota/internal/user/domain"
	"github.com/marcelofabianov/pejota/internal/user/port"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) port.UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUser(input port.GetUserInputRepository) (port.GetUserOutputRepository, error) {
	fmt.Println("Finding user with public_id:", input.PublicID)

	uuidStr := uuid.New().String()

	// Mocking user simulation SQL SELECT result
	user := domain.User{
		PublicID:     uuidStr,
		Name:         "Marcelo Fabiano",
		Email:        "marcelo@email.com",
		Role:         domain.RoleDeveloper,
		LoginEnabled: true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	output := port.GetUserOutputRepository{
		PublicID:     user.PublicID,
		Name:         user.Name,
		Email:        user.Email,
		Role:         string(user.Role),
		LoginEnabled: user.LoginEnabled,
		CreatedAt:    user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    user.UpdatedAt.Format(time.RFC3339),
	}

	return output, nil
}

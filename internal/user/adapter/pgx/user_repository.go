package pgx

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
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
	sql := `
		SELECT public_id, name, email, role, login_enabled, created_at, updated_at
		FROM users
		WHERE public_id = $1 AND deleted_at IS NULL
	`

	var user domain.User
	err := r.db.QueryRow(context.Background(), sql, input.PublicID).
		Scan(&user.PublicID, &user.Name, &user.Email, &user.Role, &user.LoginEnabled, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return port.GetUserOutputRepository{}, err
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

func (r *UserRepository) CreateUser(input port.CreateUserInputRepository) (port.CreateUserOutputRepository, error) {
	sql := `
		INSERT INTO users (public_id, name, email, password, role, login_enabled, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING public_id
	`

	var publicID string
	err := r.db.QueryRow(context.Background(), sql, input.PublicID, input.Name, input.Email, input.Password, input.Role, input.LoginEnabled, input.CreatedAt, input.UpdatedAt).
		Scan(&publicID)

	if err != nil {
		return port.CreateUserOutputRepository{}, err
	}

	output := port.CreateUserOutputRepository{
		PublicID:     publicID,
		Name:         input.Name,
		Email:        input.Email,
		LoginEnabled: input.LoginEnabled,
		Role:         input.Role,
		CreatedAt:    input.CreatedAt,
		UpdatedAt:    input.UpdatedAt,
	}

	return output, nil
}

func (r *UserRepository) FindUserIDByPublicID(input port.FindUserIDByPublicIDInputRepository) (port.FindUserIDByPublicIDOutputRepository, error) {
	sql := `SELECT id FROM users WHERE public_id = $1 AND deleted_at IS NULL`

	var user domain.User
	err := r.db.QueryRow(context.Background(), sql, input.PublicID).Scan(&user.ID)

	if err != nil {
		return port.FindUserIDByPublicIDOutputRepository{}, err
	}

	output := port.FindUserIDByPublicIDOutputRepository{
		ID: user.ID,
	}

	return output, nil
}

func (r *UserRepository) DeleteUser(input port.DeleteUserInputRepository) (port.DeleteUserOutputRepository, error) {
	sql := `UPDATE users SET deleted_at = $1 WHERE id = $2 RETURNING true`

	_, err := r.db.Exec(context.Background(), sql, input.DeletedAt, input.ID)

	if err != nil {
		return port.DeleteUserOutputRepository{}, err
	}

	output := port.DeleteUserOutputRepository{
		Success: true,
	}

	return output, nil
}

func (r *UserRepository) UpdateUser(input port.UpdateUserInputRepository) (port.UpdateUserOutputRepository, error) {
	sql := `
		UPDATE users
		SET name = $1, email = $2, role = $3, updated_at = $4
		WHERE id = $5
		RETURNING public_id, name, email, role, login_enabled, created_at, updated_at
	`

	var user domain.User
	err := r.db.QueryRow(context.Background(), sql, input.Name, input.Email, input.Role, input.UpdatedAt, input.ID).
		Scan(&user.PublicID, &user.Name, &user.Email, &user.Role, &user.LoginEnabled, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return port.UpdateUserOutputRepository{}, err
	}

	output := port.UpdateUserOutputRepository{
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

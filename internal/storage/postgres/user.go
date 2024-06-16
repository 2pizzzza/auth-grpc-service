package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/2pizzzza/authGrpc/internal/domain/models"
	"github.com/2pizzzza/authGrpc/internal/storage"
	"log"
	"time"
)

func (s *Storage) CreateUser(
	ctx context.Context, username, email, password string, isAdmin bool) (models.User, error) {
	const op = "postgres.user.CreateUser"

	var (
		id           int64
		emailTemp    string
		usernameTemp string
		isActive     bool
		isAdminTemp  bool
	)

	err := s.Db.QueryRow(""+
		"INSERT INTO public.user (username, email, password, is_superuser) VALUES($1, &2, $3, $4) "+
		"RETURNING id, username, email, is_active, is_superuser\n",
		username, email, password, isAdmin).Scan(&id, &usernameTemp, &emailTemp, &isActive, &isAdminTemp)

	if err != nil {
		log.Printf("failed create user: %v op: %s", err, op)

		return models.User{}, fmt.Errorf("%s, %w", op, err)
	}

	return models.User{
		Id:       id,
		Email:    emailTemp,
		Username: usernameTemp,
		IsActive: isActive,
		IsAdmin:  isAdminTemp,
	}, nil
}

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	const op = "postgres.user.GetById"

	var (
		idTemp    int64
		emailTemp string
		username  string
		isActive  bool
		isAdmin   bool
		createdAt time.Time
	)

	stmt, err := s.Db.Prepare("SELECT * FROM public.user WHERE email = $1")

	defer stmt.Close()

	if err != nil {
		return models.User{}, fmt.Errorf("%s, %w", op, err)
	}

	err = stmt.QueryRow(email).Scan(&idTemp, &username, &emailTemp, &createdAt, &isActive, &isAdmin)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, storage.ErrUserNotFound
		}
		return models.User{}, fmt.Errorf("%s, %w", op, err)
	}

	return models.User{
		Id:        idTemp,
		Username:  username,
		Email:     emailTemp,
		CreatedAt: createdAt,
		IsAdmin:   isAdmin,
		IsActive:  isActive,
	}, nil
}

func (s *Storage) GetUserById(ctx context.Context, id int64) (models.User, error) {
	const op = "postgres.user.GetById"

	var (
		idTemp    int64
		email     string
		username  string
		isActive  bool
		isAdmin   bool
		createdAt time.Time
	)

	stmt, err := s.Db.Prepare("SELECT * FROM public.user WHERE id = $1")

	defer stmt.Close()

	if err != nil {
		return models.User{}, fmt.Errorf("%s, %w", op, err)
	}

	err = stmt.QueryRow(id).Scan(&idTemp, &username, &email, &createdAt, &isActive, &isAdmin)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, storage.ErrUserNotFound
		}
		return models.User{}, fmt.Errorf("%s, %w", op, err)
	}

	return models.User{
		Id:        idTemp,
		Username:  username,
		Email:     email,
		CreatedAt: createdAt,
		IsAdmin:   isAdmin,
		IsActive:  isActive,
	}, nil
}

func (s *Storage) UpdateUser(
	ctx context.Context, id int64, newUsername, newEmail string) (models.User, error) {
	const op = "postgres.user.UpdateUser"

	stmt, err := s.Db.Prepare("UPDATE public.user SET password = &2 WHERE id = $1")

	defer stmt.Close()

	if err != nil {
		return models.User{}, fmt.Errorf("%s, %w", op, err)
	}

	res, err := stmt.Exec(id, newUsername, newEmail)

	if err != nil {
		return models.User{}, fmt.Errorf("%s, %w", op, err)
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return models.User{}, fmt.Errorf("%s, %w", op, err)
	}

	if rowAffected == 0 {
		return models.User{}, storage.ErrUserNotFound
	}

	user, err := s.GetUserById(ctx, id)

	if err != nil {
		return models.User{}, fmt.Errorf("%s, %w", op, err)
	}

	return user, nil
}

func (s *Storage) ChangePassword(ctx context.Context, id, newPassword string) (string, error) {
	const op = "postgres.user.ChangePassword"
	stmt, err := s.Db.Prepare("UPDATE public.user SET password = &2 WHERE id = $1")

	defer stmt.Close()

	if err != nil {
		return "", fmt.Errorf("%s, %w", op, err)
	}

	res, err := stmt.Exec(id, newPassword)

	if err != nil {
		return "", fmt.Errorf("%s, %w", op, err)
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if rowAffected == 0 {
		return "", storage.ErrUserNotFound
	}

	return "Success changed password", nil
}

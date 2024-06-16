package service

import (
	"context"
	"github.com/2pizzzza/authGrpc/internal/domain/models"
	"log/slog"
	"time"
)

type AuthService struct {
	log       *slog.Logger
	auth      AuthDb
	tokenTTL  time.Duration
	jwtSecret string
}

type AuthDb interface {
	CreateUser(ctx context.Context, username, email, password string, isAdmin bool) (models.User, error)
	GetUserById(ctx context.Context, id int64) (models.User, error)
	UpdateUser(ctx context.Context, id int64, newUsername, newEmail string) (models.User, error)
	ChangePassword(ctx context.Context, id, newPassword string) (string, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

func New(
	log *slog.Logger,
	auth AuthDb,
	tokenTTL time.Duration,
	jwtSecret string,
) *AuthService {

	return &AuthService{
		log:  log,
		auth: auth,
	}
}

package service

import (
	"context"
	"fmt"
	"github.com/2pizzzza/authGrpc/internal/lib/jwt"
	"github.com/2pizzzza/authGrpc/internal/lib/logger/sl"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func (a *AuthService) Register(ctx context.Context, email, username, password string) (string, error) {
	const op = "service.RegisterUser"

	log := a.log.With(
		slog.String("op: ", op),
	)
	log.Info("register user")

	passwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed password hash", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	user, err := a.auth.CreateUser(ctx, username, email, string(passwdHash))

	if err != nil {
		log.Warn("Failed user register")
		return "", fmt.Errorf("%s, %w", op, err)
	}

	_ = user

	accessToken, err := jwt.NewToken(user, a.jwtSecret, a.tokenTTL)
	if err != nil {
		log.Error("failed to generate token", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user registered")

	return accessToken, nil
}

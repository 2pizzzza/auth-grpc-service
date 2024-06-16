package service

import (
	"context"
	"fmt"
	"github.com/2pizzzza/authGrpc/internal/lib/jwt"
	"github.com/2pizzzza/authGrpc/internal/lib/logger/sl"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func (a *AuthService) Login(ctx context.Context, email, password string) (accessToken string, err error) {
	const op = "service.login.Login"

	log := a.log.With(
		slog.String("op: ", op),
	)
	log.Info("login user")

	user, err := a.auth.GetUserByEmail(ctx, email)
	if err != nil {
		log.Warn("User not found")
		return "", fmt.Errorf("%s, %w", op, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Info("invalid credentials", sl.Err(err))
		return "", fmt.Errorf("%s, %w", op, err)
	}

	log.Info("Login successfully")

	accessToken, err = jwt.NewToken(user, a.jwtSecret, a.tokenTTL)
	if err != nil {
		log.Warn("failed create token")
		return "", fmt.Errorf("%s, %w", op, err)
	}

	return accessToken, nil
}

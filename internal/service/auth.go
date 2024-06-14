package service

import "log/slog"

type AuthService struct {
	log  *slog.Logger
	auth AuthDb
}

type AuthDb interface {
}

func New(
	log *slog.Logger,
	auth AuthDb,
) *AuthService {

	return &AuthService{
		log:  log,
		auth: auth,
	}
}

package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Env      string
	DBConn   DatabaseConfig
	GrpcConn GrpcConfig
	JwtConn  JwtConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type GrpcConfig struct {
	GrpcPort int
	GrpcHost string
}

type JwtConfig struct {
	TokenTTL  time.Duration
	JwtSecret string
}

func MustLoad() (db *Config, err error) {
	err = godotenv.Load()

	if err != nil {
		fmt.Println("Error is occurred  on .env file please check", err)
		return nil, err
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	env := os.Getenv("ENV")

	grpcHost := os.Getenv("GRPC_HOST")
	grpcPort, _ := strconv.Atoi(os.Getenv("GRPC_PORT"))

	jwtSecret := os.Getenv("JWT_SECRET")
	tokenTTLStr := os.Getenv("TOKEN_TTl")

	tokenTTL, err := time.ParseDuration(tokenTTLStr)
	if err != nil {
		panic(fmt.Sprintf("failed to parse TokenTLL: %s", err))
	}

	return &Config{
		Env: env,
		DBConn: DatabaseConfig{
			Host:     host,
			Port:     port,
			Database: dbname,
			Username: pass,
			Password: pass,
		},
		GrpcConn: GrpcConfig{
			GrpcHost: grpcHost,
			GrpcPort: grpcPort,
		},
		JwtConn: JwtConfig{
			JwtSecret: jwtSecret,
			TokenTTL:  tokenTTL,
		},
	}, nil
}

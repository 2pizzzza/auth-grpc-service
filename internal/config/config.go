package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Env      string
	GrpcPort int
	GrpcHost string
}

func NewConfig() (db *Config, err error) {
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
	log.Printf("Succses load env %s", pass)

	return &Config{
		Host:     host,
		Port:     port,
		Database: dbname,
		Username: pass,
		Password: pass,
		Env:      env,
		GrpcPort: grpcPort,
		GrpcHost: grpcHost,
	}, nil
}

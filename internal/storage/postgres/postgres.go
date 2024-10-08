package postgres

import (
	"database/sql"
	"fmt"
	"github.com/2pizzzza/authGrpc/internal/config"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	Db *sql.DB
}

func New(con *config.Config) (*Storage, error) {

	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		con.DBConn.Host, con.DBConn.Port, con.DBConn.Username, con.DBConn.Database, con.DBConn.Password)

	connDb, err := sql.Open("postgres", conn)

	if err != nil {
		log.Printf("Error connection db: %s", err)
		return nil, err
	}
	log.Printf("Succses connect database")

	return &Storage{
		Db: connDb,
	}, nil
}


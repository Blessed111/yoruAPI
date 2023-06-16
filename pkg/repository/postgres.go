package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const teamtable = "Teams"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Sslmode  string
}

func NewPostgresDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.Sslmode))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func NewPostgresDb1() (*sqlx.DB, error) {

	username := "postgres"
	password := "55313104"
	host := "localhost"
	port := "5432"
	dbname := "YORU"
	sslmode := "disable"
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, username, dbname, password, sslmode))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

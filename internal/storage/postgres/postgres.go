package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "db"
	port     = "5432"
	user     = "latte"
	password = "latte"
	dbname   = "frappuccino"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres() *Postgres {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	log.Println("Successfully connected to the database!")

	return &Postgres{DB: db}
}

func (p *Postgres) Close() {
	if err := p.DB.Close(); err != nil {
		log.Printf("Error closing the database connection: %v", err)
	}
}

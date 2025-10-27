package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jegasape/spirex/internal/entity"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type TransactionRepository interface {
	Add(entity.Detail) error
	Edit(entity.Detail) error
	Delete(entity.Detail) error
	FindAll() []entity.Detail
}

type DbRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &DbRepository{
		DB: db,
	}
}

func (r *DbRepository) Add(e entity.Detail) error {
	return nil
}

func (r *DbRepository) Edit(e entity.Detail) error {
	return nil
}

func (r *DbRepository) Delete(e entity.Detail) error {
	return nil
}

func (r *DbRepository) FindAll() []entity.Detail {
	return nil
}

func Connection() (*sql.DB, error) {
	if os.Getenv("DB_HOST") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Printf("Warning: No .env file found, using environment variables")
		}
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var db *sql.DB
	var err error

	retries := 10
	for i := range retries {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Printf("Attempt %d: Error connecting to database: %v", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Attempt %d: Database not ready yet: %v", i+1, err)
			db.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		log.Printf("Successfully connected to database on attempt %d!", i+1)

		return db, nil
	}

	return nil, fmt.Errorf("Failed to connect to database after %d attempts: %v", retries, err)
}

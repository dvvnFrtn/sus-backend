package main

import (
	"log"
	"sus-backend/config"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/repository"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := config.SetupDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	queries := sqlc.New(db)
	repo := repository.NewSeederRepository(queries)

	categories := []sqlc.AddCategoryParams{
		{ID: uuid.New().String(), CategoryName: "Product Design"},
		{ID: uuid.New().String(), CategoryName: "Product Management"},
		{ID: uuid.New().String(), CategoryName: "Front-End Development"},
		{ID: uuid.New().String(), CategoryName: "Back-End Development"},
		{ID: uuid.New().String(), CategoryName: "Data Science"},
	}

	for _, category := range categories {
		_, err = repo.AddCategory(category)
		if err != nil {
			log.Fatalf("Failed adding category: %v", err)
		}
	}
}

package main

import (
	"database/sql"
	"fmt"
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

	// seed 5 user with user role
	for i := 0; i < 5; i++ {
		repo.AddUser(sqlc.AddUserParams{
			ID:    uuid.New().String(),
			Email: fmt.Sprintf("user.dummy%d@gmail.com", i),
			Password: sql.NullString{
				String: "$2a$10$OkddtETs010l0lUa6RhU0ugKzzaWGUTyblynVc1IJVvDKOoLfisX.",
				Valid:  true,
			},
			Name: fmt.Sprintf("User Dummy %d", i),
			Role: "user",
		})
	}

	// seed 10 user with organizer role
	for i := 0; i < 10; i++ {
		userID := uuid.New().String()
		repo.AddUser(sqlc.AddUserParams{
			ID:    userID,
			Email: fmt.Sprintf("organizer.dummy%d@gmail.com", i),
			Password: sql.NullString{
				String: "$2a$10$OkddtETs010l0lUa6RhU0ugKzzaWGUTyblynVc1IJVvDKOoLfisX.",
				Valid:  true,
			},
			Name: fmt.Sprintf("Organizer Dummy %d", i),
			Role: "organization",
		})

		// seed organization belongs to organizer
		orgID := uuid.New().String()
		repo.AddOrganization(sqlc.AddOrganizationParams{
			ID:          orgID,
			UserID:      userID,
			Name:        fmt.Sprintf("Organization Organizer %d", i),
			Description: "Posuere potenti consequat lorem, consectetur mattis et gravida. Quisque eu nullam; faucibus imperdiet elementum sapien.",
			ProfileImg: sql.NullString{
				String: "http://img.dummy.co/dummy-profile-organization.img",
				Valid:  true,
			},
		})

		// seed 5 post for each organizations
		for j := 0; j < 5; j++ {
			repo.AddPost(sqlc.AddPostParams{
				ID:             uuid.New().String(),
				OrganizationID: orgID,
				Content:        "Lorem ipsum odor amet, consectetuer adipiscing elit. Faucibus dignissim cursus vestibulum posuere maximus torquent fusce. Torquent varius cubilia fermentum quis fames proin habitasse curabitur tortor. Ultrices accumsan ultrices dictum magnis congue tincidunt etiam ad. Aptent massa faucibus urna ad dolor. Auctor eu taciti aenean, interdum purus ante. ",
				ImageContent: sql.NullString{
					String: "http://img.dummy.co/dummy-post.img",
					Valid:  true,
				},
			})
		}
	}
}

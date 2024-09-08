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

	groups := []string{"College", "Work", "Community"}

	for _, group := range groups {
		exist, err := repo.CategoryGroupExists(group)
		if err != nil {
			log.Fatalf("Failed checking category group: %v", err)
		}

		if exist < 1 {
			_, err := repo.AddCategoryGroup(group)
			if err != nil {
				log.Fatalf("Failed adding category group: %v", err)
			}
		}
	}

	collegeGroupID := GetGroupID(repo, "College")
	workGroupID := GetGroupID(repo, "Work")
	communityGroupID := GetGroupID(repo, "Community")

	categories := []sqlc.AddCategoryParams{
		{ID: uuid.New().String(), CategoryName: "Academic Societies", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Self-Development Organizations", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Cultural Organizations", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Recreational/Sports Clubs", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Professional Organizations", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Service Organizations", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Student Government", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Artistic/Music Groups", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Social Organizations", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Environmental Organizations", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Religious/Spiritual Organizations", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Political Organizations", GroupID: collegeGroupID},
		{ID: uuid.New().String(), CategoryName: "Health and Wellness Groups", GroupID: collegeGroupID},

		{ID: uuid.New().String(), CategoryName: "Professional Development Organizations", GroupID: workGroupID},
		{ID: uuid.New().String(), CategoryName: "Employee Resource Groups (ERGs)", GroupID: workGroupID},
		{ID: uuid.New().String(), CategoryName: "Wellness Committees", GroupID: workGroupID},
		{ID: uuid.New().String(), CategoryName: "Social Committees", GroupID: workGroupID},
		{ID: uuid.New().String(), CategoryName: "Innovation Hubs", GroupID: workGroupID},
		{ID: uuid.New().String(), CategoryName: "Sustainability Committees", GroupID: workGroupID},
		{ID: uuid.New().String(), CategoryName: "Union/Workers Councils", GroupID: workGroupID},
		{ID: uuid.New().String(), CategoryName: "Charity/Community Outreach Groups", GroupID: workGroupID},
		{ID: uuid.New().String(), CategoryName: "Networking Groups", GroupID: workGroupID},
		{ID: uuid.New().String(), CategoryName: "Mentorship Programs", GroupID: workGroupID},

		{ID: uuid.New().String(), CategoryName: "Self-Development Groups", GroupID: communityGroupID},
		{ID: uuid.New().String(), CategoryName: "Social Clubs", GroupID: communityGroupID},
		{ID: uuid.New().String(), CategoryName: "Recreational/Sports Leagues", GroupID: communityGroupID},
		{ID: uuid.New().String(), CategoryName: "Environmental Groups", GroupID: communityGroupID},
		{ID: uuid.New().String(), CategoryName: "Volunteer/Service Organizations", GroupID: communityGroupID},
		{ID: uuid.New().String(), CategoryName: "Political Action Groups", GroupID: communityGroupID},
		{ID: uuid.New().String(), CategoryName: "Cultural Groups", GroupID: communityGroupID},
		{ID: uuid.New().String(), CategoryName: "Artistic/Music Groups", GroupID: communityGroupID},
		{ID: uuid.New().String(), CategoryName: "Business Networks", GroupID: communityGroupID},
		{ID: uuid.New().String(), CategoryName: "Health and Wellness Groups", GroupID: communityGroupID},
		{ID: uuid.New().String(), CategoryName: "Educational Organizations", GroupID: communityGroupID},
	}

	for _, category := range categories {
		exist, err := repo.CategoryExists(sqlc.CategoryExistsParams{
			CategoryName: category.CategoryName,
			GroupID:      category.GroupID,
		})
		if err != nil {
			log.Fatalf("Failed checking category: %v", err)
		}

		if exist < 1 {
			_, err = repo.AddCategory(category)
			if err != nil {
				log.Fatalf("Failed adding category: %v", err)
			}
		}
	}
}

func GetGroupID(r repository.SeederRepository, s string) int32 {
	id, err := r.GetGroupIDByName(s)
	if err != nil {
		log.Fatalf("Failed getting category group: %v", err)
	}
	return id
}

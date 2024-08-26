package dto

import (
	"sus-backend/internal/db/sqlc"
	"time"
)

type OrganizationCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type OrganizationUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type OrganizationResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ToOrganizationResponse(organization *sqlc.Organization) *OrganizationResponse {
	return &OrganizationResponse{
		ID:          organization.ID,
		Name:        organization.Name,
		Description: organization.Description,
		CreatedAt:   organization.CreatedAt.Time,
		UpdatedAt:   organization.UpdatedAt.Time,
	}
}

func ToOrganizationResponses(organizations *[]sqlc.Organization) []OrganizationResponse {
	organizationResponses := []OrganizationResponse{}
	for _, organization := range *organizations {
		organizationResponse := OrganizationResponse{
			ID:          organization.ID,
			Name:        organization.Name,
			Description: organization.Description,
			CreatedAt:   organization.CreatedAt.Time,
			UpdatedAt:   organization.UpdatedAt.Time,
		}

		organizationResponses = append(organizationResponses, organizationResponse)
	}

	return organizationResponses
}

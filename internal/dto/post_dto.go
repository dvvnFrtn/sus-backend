package dto

import (
	"sus-backend/internal/db/sqlc"
	"time"
)

type PostCreateRequest struct {
	Content      string `json:"content" binding:"required"`
	ImageContent string `json:"imageContent"`
}

type PostResponse struct {
	ID           string            `json:"id"`
	Content      string            `json:"content"`
	ImageContent string            `json:"imageContent,omitempty"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
	Organization *WithOrganization `json:"organization,omitempty"`
}

type WithOrganization struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ProfileImg string `json:"profile_img,omitempty"`
}

func ToPostResponse(post *sqlc.FindPostByIdRow) *PostResponse {
	return &PostResponse{
		ID:           post.ID,
		Content:      post.Content,
		ImageContent: post.ImageContent.String,
		CreatedAt:    post.CreatedAt.Time,
		UpdatedAt:    post.UpdatedAt.Time,
		Organization: &WithOrganization{
			ID:         post.OrganizationID,
			Name:       post.Name,
			ProfileImg: post.ProfileImg.String,
		},
	}
}

func ToPostResponses(posts *[]sqlc.FindPostByIdRow) []PostResponse {
	postResponses := []PostResponse{}
	for _, post := range *posts {
		postResponse := PostResponse{
			ID:           post.ID,
			Content:      post.Content,
			ImageContent: post.ImageContent.String,
			CreatedAt:    post.CreatedAt.Time,
			UpdatedAt:    post.UpdatedAt.Time,
			Organization: &WithOrganization{
				ID:         post.OrganizationID,
				Name:       post.Name,
				ProfileImg: post.ProfileImg.String,
			},
		}

		postResponses = append(postResponses, postResponse)
	}

	return postResponses
}

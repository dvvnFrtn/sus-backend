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
	ID           string    `json:"id"`
	Content      string    `json:"content"`
	ImageContent string    `json:"imageContent"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func ToPostResponse(post *sqlc.Post) *PostResponse {
	return &PostResponse{
		ID:           post.ID,
		Content:      post.Content,
		ImageContent: post.ImageContent.String,
		CreatedAt:    post.CreatedAt.Time,
		UpdatedAt:    post.UpdatedAt.Time,
	}
}

func ToPostResponses(posts *[]sqlc.Post) []PostResponse {
	postResponses := []PostResponse{}
	for _, post := range *posts {
		postResponse := PostResponse{
			ID:           post.ID,
			Content:      post.Content,
			ImageContent: post.ImageContent.String,
			CreatedAt:    post.CreatedAt.Time,
			UpdatedAt:    post.UpdatedAt.Time,
		}

		postResponses = append(postResponses, postResponse)
	}

	return postResponses
}

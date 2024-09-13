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
	ImageContent string            `json:"image_content,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Organization *WithOrganization `json:"organization,omitempty"`
	Likes        int               `json:"likes"`
}

type PostLikesResponse struct {
	UserID     string    `json:"user_id"`
	Username   string    `json:"username"`
	ProfileImg string    `json:"profile_img"`
	LikedAt    time.Time `json:"liked_at"`
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
		Likes: int(post.Likes),
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
			Likes: int(post.Likes),
		}

		postResponses = append(postResponses, postResponse)
	}

	return postResponses
}

func ToPostLikesResponse(postLikes *[]sqlc.FindPostLikesRow) []PostLikesResponse {
	postLikesResponses := []PostLikesResponse{}
	for _, pl := range *postLikes {
		postLikeResponse := PostLikesResponse{
			UserID:     pl.UserID,
			Username:   pl.Name,
			ProfileImg: pl.Img.String,
			LikedAt:    pl.LikedAt.Time,
		}

		postLikesResponses = append(postLikesResponses, postLikeResponse)
	}

	return postLikesResponses
}

package dto

import (
	"sus-backend/internal/db/sqlc"
	"time"
)

type PostCreateRequest struct {
	Content      string `json:"content" binding:"required"`
	ImageContent string `json:"imageContent"`
}

type CommentPostRequest struct {
	PostID  string `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostResponse struct {
	ID           string            `json:"id"`
	Content      string            `json:"content"`
	ImageContent string            `json:"image_content"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Likes        int               `json:"likes_count"`
	Comments     int               `json:"comments_count"`
	Organization *withOrganization `json:"organization"`
}

type PostLikesResponse struct {
	LikedAt time.Time `json:"liked_at"`
	User    *withUser `json:"user"`
}

type PostCommentsResponse struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	User      *withUser `json:"user"`
}

func ToPostResponse(post *sqlc.FindPostByIdRow) *PostResponse {
	return &PostResponse{
		ID:           post.ID,
		Content:      post.Content,
		ImageContent: post.ImageContent.String,
		CreatedAt:    post.CreatedAt.Time,
		UpdatedAt:    post.UpdatedAt.Time,
		Organization: &withOrganization{
			ID:         post.OrganizationID,
			Name:       post.Name,
			ProfileImg: post.ProfileImg.String,
		},
		Likes:    int(post.LikeCount),
		Comments: int(post.CommentCount),
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
			Organization: &withOrganization{
				ID:         post.OrganizationID,
				Name:       post.Name,
				ProfileImg: post.ProfileImg.String,
			},
			Likes:    int(post.LikeCount),
			Comments: int(post.CommentCount),
		}

		postResponses = append(postResponses, postResponse)
	}

	return postResponses
}

func ToPostLikesResponse(postLikes *[]sqlc.FindPostLikesRow) []PostLikesResponse {
	postLikesResponses := []PostLikesResponse{}
	for _, pl := range *postLikes {
		postLikeResponse := PostLikesResponse{
			User: &withUser{
				ID:         pl.UserID,
				Username:   pl.Name,
				ProfileImg: pl.Img.String,
			},
			LikedAt: pl.LikedAt.Time,
		}

		postLikesResponses = append(postLikesResponses, postLikeResponse)
	}

	return postLikesResponses
}

func ToPostCommentsResponse(postComments *[]sqlc.FindPostCommentsRow) []PostCommentsResponse {
	postCommentsResponses := []PostCommentsResponse{}
	for _, pc := range *postComments {
		postCommentResponse := PostCommentsResponse{
			ID: pc.ID,
			User: &withUser{
				ID:         pc.UserID,
				Username:   pc.Name,
				ProfileImg: pc.Img.String,
			},
			Content:   pc.Content,
			CreatedAt: pc.CreatedAt.Time,
		}

		postCommentsResponses = append(postCommentsResponses, postCommentResponse)
	}

	return postCommentsResponses
}

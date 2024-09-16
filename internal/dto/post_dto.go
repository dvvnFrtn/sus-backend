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
	ImageContent string            `json:"image_content,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Organization *WithOrganization `json:"organization,omitempty"`
	Likes        int               `json:"likes"`
	Comments     int               `json:"comments"`
}

type PostLikesResponse struct {
	UserID     string    `json:"user_id"`
	Username   string    `json:"username"`
	ProfileImg string    `json:"profile_img"`
	LikedAt    time.Time `json:"liked_at"`
}

type PostCommentsResponse struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Username   string    `json:"username"`
	ProfileImg string    `json:"profile_img"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
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
			Organization: &WithOrganization{
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
			UserID:     pl.UserID,
			Username:   pl.Name,
			ProfileImg: pl.Img.String,
			LikedAt:    pl.LikedAt.Time,
		}

		postLikesResponses = append(postLikesResponses, postLikeResponse)
	}

	return postLikesResponses
}

func ToPostCommentsResponse(postComments *[]sqlc.FindPostCommentsRow) []PostCommentsResponse {
	postCommentsResponses := []PostCommentsResponse{}
	for _, pc := range *postComments {
		postCommentResponse := PostCommentsResponse{
			ID:         pc.ID,
			UserID:     pc.UserID,
			Username:   pc.Name,
			ProfileImg: pc.Img.String,
			Content:    pc.Content,
			CreatedAt:  pc.CreatedAt.Time,
		}

		postCommentsResponses = append(postCommentsResponses, postCommentResponse)
	}

	return postCommentsResponses
}

package service

import (
	"database/sql"
	"errors"
	"fmt"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/dto"
	"sus-backend/internal/repository"
	_error "sus-backend/pkg/err"

	"github.com/google/uuid"
)

type PostService interface {
	CreatePost(string, dto.PostCreateRequest) (*dto.ResponseID, error)
	GetPostById(string) (*dto.PostResponse, error)
	GetPostsByOrganization(string) ([]dto.PostResponse, error)
	GetAllPosts() ([]dto.PostResponse, error)
	DeletePost(string, string) error
	LikedPost(string, string) error
	UnlikedPost(string, string) error
	GetPostLikes(string) ([]dto.PostLikesResponse, error)
	CommentPost(string, dto.CommentPostRequest) (*dto.ResponseID, error)
	GetPostComments(string) ([]dto.PostCommentsResponse, error)
	DeleteComment(string, string) error
}

type postService struct {
	repo    repository.PostRepository
	orgRepo repository.OrganizationRepository
}

func NewPostService(repo repository.PostRepository, orgRepo repository.OrganizationRepository) PostService {
	return &postService{repo, orgRepo}
}

func (s *postService) CreatePost(organizationID string, req dto.PostCreateRequest) (*dto.ResponseID, error) {
	// Error: there is no organization associated
	if organizationID == "" {
		return nil, _error.ErrNoOrganization
	}

	params := sqlc.AddPostParams{
		ID:             uuid.New().String(),
		OrganizationID: organizationID,
		Content:        req.Content,
		ImageContent:   sql.NullString{String: req.ImageContent, Valid: req.ImageContent != ""},
	}

	if _, err := s.repo.Create(params); err != nil {
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	return dto.NewResponseID(params.ID), nil
}

func (s *postService) GetPostById(postID string) (*dto.PostResponse, error) {
	post, err := s.repo.FindById(postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, _error.ErrNotFound
		}
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	return dto.ToPostResponse(&post), nil
}

func (s *postService) GetPostsByOrganization(organizationID string) ([]dto.PostResponse, error) {
	if _, err := s.orgRepo.FindById(organizationID); err != nil {
		return nil, _error.ErrNotFound
	}

	posts, err := s.repo.FindByOrganization(organizationID)
	if err != nil {
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	return dto.ToPostResponses(&posts), nil
}

// TODO: implement using timeline, followed by user
func (s *postService) GetAllPosts() ([]dto.PostResponse, error) {
	posts, err := s.repo.ListAll()
	if err != nil {
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	return dto.ToPostResponses(&posts), nil
}

func (s *postService) DeletePost(organizationID string, postID string) error {
	post, err := s.repo.FindById(postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return _error.ErrNoDeleted
		}
		fmt.Println(err)
		return _error.ErrInternal
	}

	// Error: resource not belongs to organizer
	if post.OrganizationID != organizationID {
		return _error.ErrForbidden
	}

	if err := s.repo.Delete(postID); err != nil {
		fmt.Println(err)
		return _error.ErrInternal
	}

	return nil
}

func (s *postService) LikedPost(authID string, postID string) error {
	if _, err := s.repo.FindById(postID); err != nil {
		return _error.ErrNotFound
	}

	params := sqlc.LikedPostParams{
		UserID: authID,
		PostID: postID,
	}

	count, err := s.repo.IsLiked(sqlc.IsLikedParams(params))
	if err != nil {
		fmt.Println(err)
		return _error.ErrInternal
	}

	// Error: user already liked the post
	if count >= 1 {
		return _error.ErrAlreadyLiked
	}

	if _, err := s.repo.LikedPost(params); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *postService) UnlikedPost(authID string, postID string) error {
	if _, err := s.repo.FindById(postID); err != nil {
		return _error.ErrNotFound
	}

	params := sqlc.UnlikedPostParams{
		UserID: authID,
		PostID: postID,
	}

	count, err := s.repo.IsLiked(sqlc.IsLikedParams(params))
	if err != nil {
		fmt.Println(err)
		return _error.ErrInternal
	}

	// Error: user not liked the post yet
	if count <= 0 {
		return _error.ErrNotLiked
	}

	if err := s.repo.UnlikedPost(params); err != nil {
		fmt.Println(err)
		return _error.ErrInternal
	}

	return nil
}

func (s *postService) GetPostLikes(postID string) ([]dto.PostLikesResponse, error) {
	if _, err := s.repo.FindById(postID); err != nil {
		return nil, _error.ErrNotFound
	}

	postLikes, err := s.repo.FindPostLikes(postID)
	if err != nil {
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	return dto.ToPostLikesResponse(&postLikes), nil
}

func (s *postService) CommentPost(authID string, req dto.CommentPostRequest) (*dto.ResponseID, error) {
	if _, err := s.repo.FindById(req.PostID); err != nil {
		return nil, _error.ErrNotFound
	}

	params := sqlc.CommentPostParams{
		ID:      uuid.New().String(),
		UserID:  authID,
		PostID:  req.PostID,
		Content: req.Content,
	}

	if _, err := s.repo.CommentPost(params); err != nil {
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	return dto.NewResponseID(params.ID), nil
}

func (s *postService) GetPostComments(postID string) ([]dto.PostCommentsResponse, error) {
	if _, err := s.repo.FindById(postID); err != nil {
		return nil, _error.ErrNotFound
	}

	postComments, err := s.repo.FindPostComments(postID)
	if err != nil {
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	return dto.ToPostCommentsResponse(&postComments), nil
}

func (s *postService) DeleteComment(authID string, commentID string) error {
	comment, err := s.repo.FindCommentById(commentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return _error.ErrNoDeleted
		}
		fmt.Println(err)
		return _error.ErrInternal
	}

	// Error: resource not belongs to user
	if comment.UserID != authID {
		return _error.ErrForbidden
	}

	if err := s.repo.DeleteComment(commentID); err != nil {
		fmt.Println(err)
		return _error.ErrInternal
	}

	return nil
}

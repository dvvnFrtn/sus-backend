package service

import (
	"database/sql"
	"errors"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/dto"
	"sus-backend/internal/repository"
	_error "sus-backend/pkg/err"

	"github.com/google/uuid"
)

type PostService interface {
	CreatePost(string, dto.PostCreateRequest) (*dto.ResponseID, error)
	FindPostById(string) (*dto.PostResponse, error)
	ListAllPosts() ([]dto.PostResponse, error)
	DeletePost(string, string) error
}

type postService struct {
	repo    repository.PostRepository
	orgServ OrganizationService
	orgRepo repository.OrganizationRepository
}

func NewPostService(repo repository.PostRepository, orgServ OrganizationService, orgRepo repository.OrganizationRepository) PostService {
	return &postService{repo, orgServ, orgRepo}
}

// auth: organizer
func (s *postService) CreatePost(organizationID string, req dto.PostCreateRequest) (*dto.ResponseID, error) {
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
		return nil, _error.ErrInternal
	}

	return dto.NewResponseID(params.ID), nil
}

// auth: organizer, user
func (s *postService) FindPostById(organizationID string) (*dto.PostResponse, error) {
	post, err := s.repo.FindById(organizationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, _error.ErrNotFound
		}
		return nil, _error.ErrInternal
	}

	return dto.ToPostResponse(&post), nil
}

// TODO: implement using timeline, followed by user
// auth:
func (s *postService) ListAllPosts() ([]dto.PostResponse, error) {
	posts, err := s.repo.ListAll()
	if err != nil {
		return nil, _error.ErrInternal
	}

	return dto.ToPostResponses(&posts), nil
}

// auth: organizer
func (s *postService) DeletePost(organizationID string, postID string) error {
	post, err := s.repo.FindById(postID)
	if err != nil {
		return _error.ErrNoDeleted
	}

	if post.OrganizationID != organizationID {
		return _error.ErrForbidden
	}

	if err := s.repo.Delete(postID); err != nil {
		return _error.ErrInternal
	}

	return nil
}

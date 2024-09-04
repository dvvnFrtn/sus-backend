package service

import (
	"database/sql"
	"errors"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/dto"
	"sus-backend/internal/repository"
	"time"

	"github.com/google/uuid"
)

type PostService interface {
	CreatePost(string, dto.PostCreateRequest) (*dto.ResponseID, error)
	FindPostById(string) (*dto.PostResponse, error)
	FindPostByOrganization(string) ([]dto.PostResponse, error)
	ListAllPosts() ([]dto.PostResponse, error)
	DeletePost(id string) error
}

type postService struct {
	repo    repository.PostRepository
	orgServ OrganizationService
}

func NewPostService(repo repository.PostRepository, orgServ OrganizationService) PostService {
	return &postService{repo, orgServ}
}

func (s *postService) CreatePost(orgId string, req dto.PostCreateRequest) (*dto.ResponseID, error) {
	post := sqlc.AddPostParams{
		ID:             uuid.New().String(),
		OrganizationID: orgId, // TODO: handle foreign key constraint
		Content:        req.Content,
		ImageContent:   sql.NullString{String: req.ImageContent, Valid: req.ImageContent != ""},
		CreatedAt:      sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:      sql.NullTime{Time: time.Now(), Valid: true},
	}

	_, err := s.repo.Create(post)
	if err != nil {
		return nil, err
	}

	return dto.NewResponseID(post.ID), nil
}

func (s *postService) FindPostById(id string) (*dto.PostResponse, error) {
	post, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("resource_not_found")
		}
		return nil, err
	}

	return dto.ToPostResponse(&post), nil
}

func (s *postService) FindPostByOrganization(id string) ([]dto.PostResponse, error) {
	_, err := s.orgServ.FindOrganizationById(id)
	if err != nil {
		return nil, errors.New("associated_organization_resource_not_found")
	}
	posts, err := s.repo.FindByOrganization(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("resource_not_found")
		}
		return nil, err
	}

	return dto.ToPostResponses(&posts), nil
}

func (s *postService) ListAllPosts() ([]dto.PostResponse, error) {
	posts, err := s.repo.ListAll()
	if err != nil {
		return nil, err
	}

	return dto.ToPostResponses(&posts), nil
}

func (s *postService) DeletePost(id string) error {
	_, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("no_resource_to_delete")
		}
		return err
	}

	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

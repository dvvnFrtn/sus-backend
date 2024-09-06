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

type OrganizationService interface {
	CreateOrganization(string, dto.OrganizationCreateRequest) (*dto.ResponseID, error)
	FindOrganizationById(string) (*dto.OrganizationResponse, error)
	ListAllOrganizations() ([]dto.OrganizationResponse, error)
	UpdateOrganization(string, string, dto.OrganizationUpdateRequest) (*dto.ResponseID, error)
	DeleteOrganization(string, string) error
}

type organizationService struct {
	repo repository.OrganizationRepository
}

func NewOrganizationService(repo repository.OrganizationRepository) OrganizationService {
	return &organizationService{repo}
}

func (s *organizationService) CreateOrganization(authID string, req dto.OrganizationCreateRequest) (*dto.ResponseID, error) {
	_, err := s.repo.FindByUserId(authID)
	if err == nil {
		return nil, _error.ErrConflict
	}

	params := sqlc.AddOrganizationParams{
		ID:          uuid.New().String(),
		UserID:      authID,
		Name:        req.Name,
		Description: req.Description,
		HeaderImg:   sql.NullString{},
		ProfileImg:  sql.NullString{},
	}

	_, err = s.repo.Create(params)
	if err != nil {
		return nil, _error.ErrInternal
	}

	return dto.NewResponseID(params.ID), nil
}

func (s *organizationService) FindOrganizationById(id string) (*dto.OrganizationResponse, error) {
	organization, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, _error.ErrNotFound
		}
		return nil, _error.ErrInternal
	}

	return dto.ToOrganizationResponse(&organization), nil
}

func (s *organizationService) ListAllOrganizations() ([]dto.OrganizationResponse, error) {
	organizations, err := s.repo.ListAll()
	if err != nil {
		return nil, _error.ErrInternal
	}

	return dto.ToOrganizationResponses(&organizations), nil
}

func (s *organizationService) UpdateOrganization(authID string, organizationID string, req dto.OrganizationUpdateRequest) (*dto.ResponseID, error) {
	organization, err := s.repo.FindById(organizationID)
	if err != nil {
		return nil, _error.ErrNoUpdated
	}

	if organization.UserID != authID {
		return nil, _error.ErrForbidden
	}

	params := &sqlc.UpdateOrganizationParams{
		Name:        req.Name,
		Description: req.Description,
		HeaderImg:   sql.NullString{Valid: false},
		ProfileImg:  sql.NullString{Valid: false},
		ID:          organizationID,
	}

	if _, err := s.repo.Update(*params); err != nil {
		return nil, _error.ErrInternal
	}

	return dto.NewResponseID(params.ID), nil
}

func (s *organizationService) DeleteOrganization(authID string, organizationID string) error {
	organization, err := s.repo.FindById(organizationID)
	if err != nil {
		return _error.ErrNoDeleted
	}

	if organization.UserID != authID {
		return _error.ErrForbidden
	}

	if err := s.repo.Delete(organizationID); err != nil {
		return _error.ErrInternal
	}

	return nil
}

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

type OrganizationService interface {
	CreateOrganization(string, dto.OrganizationCreateRequest) (*dto.ResponseID, error)
	GetOrganizationById(string) (*dto.OrganizationResponse, error)
	GetAllOrganizations() ([]dto.OrganizationResponse, error)
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
	count, err := s.repo.IsExist(authID)
	if err != nil {
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	// Error: one organizer only can create one organization
	if count >= 1 {
		return nil, _error.ErrConflict
	}

	params := sqlc.AddOrganizationParams{
		ID:          uuid.New().String(),
		UserID:      authID,
		Name:        req.Name,
		Description: req.Description,
		// Not Implemented
		HeaderImg:  sql.NullString{},
		ProfileImg: sql.NullString{},
	}

	if _, err := s.repo.Create(params); err != nil {
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	return dto.NewResponseID(params.ID), nil
}

func (s *organizationService) GetOrganizationById(id string) (*dto.OrganizationResponse, error) {
	organization, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, _error.ErrNotFound
		}
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	return dto.ToOrganizationResponse(&organization), nil
}

func (s *organizationService) GetAllOrganizations() ([]dto.OrganizationResponse, error) {
	organizations, err := s.repo.ListAll()
	if err != nil {
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	return dto.ToOrganizationResponses(&organizations), nil
}

func (s *organizationService) UpdateOrganization(authID string, organizationID string, req dto.OrganizationUpdateRequest) (*dto.ResponseID, error) {
	organization, err := s.repo.FindById(organizationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, _error.ErrNoUpdated
		}
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	// Error: resource not belongs to authenticated user
	if organization.UserID != authID {
		return nil, _error.ErrForbidden
	}

	params := sqlc.UpdateOrganizationParams{
		Name:        req.Name,
		Description: req.Description,
		// Not Implemented
		HeaderImg:  sql.NullString{},
		ProfileImg: sql.NullString{},
		ID:         organizationID,
	}

	if _, err := s.repo.Update(params); err != nil {
		fmt.Println(err)
		return nil, _error.ErrInternal
	}

	return dto.NewResponseID(params.ID), nil
}

func (s *organizationService) DeleteOrganization(authID string, organizationID string) error {
	organization, err := s.repo.FindById(organizationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return _error.ErrNoDeleted
		}
		fmt.Println(err)
		return _error.ErrInternal
	}

	// Error: resource not belongs to authenticated user
	if organization.UserID != authID {
		return _error.ErrForbidden
	}

	if err := s.repo.Delete(organizationID); err != nil {
		fmt.Println(err)
		return _error.ErrInternal
	}

	return nil
}

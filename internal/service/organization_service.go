package service

import (
	"database/sql"
	"fmt"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/dto"
	"sus-backend/internal/repository"
	"time"

	"github.com/google/uuid"
)

type OrganizationService interface {
	CreateOrganization(dto.OrganizationCreateRequest) (*dto.ResponseID, error)
	FindOrganizationById(string) (*dto.OrganizationResponse, error)
	ListAllOrganizations() ([]dto.OrganizationResponse, error)
	UpdateOrganization(string, dto.OrganizationUpdateRequest) (*dto.ResponseID, error)
	DeleteOrganization(string) error
	GetCategories() ([]sqlc.Category, error)
}

type organizationService struct {
	repo repository.OrganizationRepository
}

func NewOrganizationService(repo repository.OrganizationRepository) OrganizationService {
	return &organizationService{repo}
}

func (s *organizationService) CreateOrganization(req dto.OrganizationCreateRequest) (*dto.ResponseID, error) {
	organization := sqlc.AddOrganizationParams{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		HeaderImg:   sql.NullString{},
		ProfileImg:  sql.NullString{},
		CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	}

	_, err := s.repo.Create(organization)
	if err != nil {
		return nil, err
	}

	return &dto.ResponseID{ID: organization.ID}, nil
}

func (s *organizationService) FindOrganizationById(id string) (*dto.OrganizationResponse, error) {
	organization, err := s.repo.FindById(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dto.ToOrganizationResponse(&organization), nil
}

func (s *organizationService) ListAllOrganizations() ([]dto.OrganizationResponse, error) {
	organizations, err := s.repo.ListAll()
	if err != nil {
		return nil, err
	}

	return dto.ToOrganizationResponses(&organizations), nil
}

func (s *organizationService) UpdateOrganization(id string, req dto.OrganizationUpdateRequest) (*dto.ResponseID, error) {
	organization, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	organizationParams := &sqlc.UpdateOrganizationParams{
		Name:        req.Name,
		Description: req.Description,
		HeaderImg:   sql.NullString{Valid: false},
		ProfileImg:  sql.NullString{Valid: false},
		ID:          organization.ID,
	}

	_, err = s.repo.Update(*organizationParams)
	if err != nil {
		return nil, err
	}

	return &dto.ResponseID{ID: organization.ID}, nil
}

func (s *organizationService) DeleteOrganization(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *organizationService) GetCategories() ([]sqlc.Category, error) {
	categories, err := s.repo.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

package repository

import (
	"context"
	"database/sql"
	"sus-backend/internal/db/sqlc"
)

type OrganizationRepository interface {
	Create(sqlc.AddOrganizationParams) (sql.Result, error)
	FindById(string) (sqlc.Organization, error)
	ListAll() ([]sqlc.Organization, error)
	Update(sqlc.UpdateOrganizationParams) (sql.Result, error)
	Delete(string) error
	FindByUserId(id string) (sqlc.Organization, error)
}

type organizationRepository struct {
	db *sqlc.Queries
}

func NewOrganizationRepository(db *sqlc.Queries) OrganizationRepository {
	return &organizationRepository{db}
}

func (r *organizationRepository) Create(in sqlc.AddOrganizationParams) (sql.Result, error) {
	return r.db.AddOrganization(context.Background(), in)
}

func (r *organizationRepository) FindById(id string) (sqlc.Organization, error) {
	org, err := r.db.FindOrganizationById(context.Background(), id)
	return org, err
}

func (r *organizationRepository) ListAll() ([]sqlc.Organization, error) {
	return r.db.ListOrganization(context.Background())
}

func (r *organizationRepository) Update(in sqlc.UpdateOrganizationParams) (sql.Result, error) {
	return r.db.UpdateOrganization(context.Background(), in)
}

func (r *organizationRepository) Delete(id string) error {
	return r.db.DeleteOrganization(context.Background(), id)
}

func (r *organizationRepository) FindByUserId(id string) (sqlc.Organization, error) {
	return r.db.FindOrganizationByUserId(context.Background(), id)
}

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

type ActivityService interface {
	GetActivityByID(string) (*dto.ActivityResponse, error)
	GetActivitiesByOrganizationID(string) ([]dto.ActivityResponse, error)
	CreateActivity(dto.ActivityCreateReq, string) (*dto.ResponseID, error)
	DeleteActivity(string, string) error
}

type activityService struct {
	repo repository.ActivityRepository
}

func NewActivityService(r repository.ActivityRepository) ActivityService {
	return &activityService{r}
}

func (s *activityService) GetActivityByID(id string) (*dto.ActivityResponse, error) {
	activity, err := s.repo.GetActivityByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("resource not found")
		}
		return nil, err
	}
	return dto.ToActivityResponse(&activity), nil
}

func (s *activityService) GetActivitiesByOrganizationID(org_id string) ([]dto.ActivityResponse, error) {
	activities, err := s.repo.GetActivitiesByOrganizationID(org_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("resources not found")
		}
		return nil, err
	}
	return dto.ToActivityResponses(&activities), nil
}

func (s *activityService) CreateActivity(arg dto.ActivityCreateReq, org_id string) (*dto.ResponseID, error) {
	startTime, err := time.Parse("2006-01-02 15:04", arg.StartTime)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse("2006-01-02 15:04", arg.EndTime)
	if err != nil {
		return nil, err
	}

	input := sqlc.CreateActivityParams{
		ID:             uuid.New().String(),
		OrganizationID: org_id,
		Title:          sql.NullString{String: arg.Title, Valid: arg.Title != ""},
		Note:           arg.Note,
		StartTime:      sql.NullTime{Time: startTime, Valid: !startTime.IsZero()},
		EndTime:        sql.NullTime{Time: endTime, Valid: !endTime.IsZero()},
	}

	_, err = s.repo.CreateActivity(input)
	if err != nil {
		return nil, err
	}

	return dto.NewResponseID(input.ID), nil
}

func (s *activityService) DeleteActivity(id string, org_id string) error {
	actv, err := s.repo.GetActivityByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("resource not found")
		}
		return err
	}

	if actv.OrganizationID != org_id {
		return errors.New("access denied - activity does not belong to user")
	}

	err = s.repo.DeleteActivity(id)
	if err != nil {
		return err
	}
	return nil
}

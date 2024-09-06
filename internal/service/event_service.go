package service

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/dto"
	"sus-backend/internal/repository"
	"time"

	"github.com/google/uuid"
)

type EventService interface {
	GetEvents() ([]dto.EventResponse, error)
	GetEventsByCategory([]string) ([]dto.EventResponse, error)
	GetEventByID(string) (*dto.EventResponse, error)
	CreateEvent(string, dto.CreateEventReq) (*dto.ResponseID, error)
	DeleteEvent(string, string) error
}

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(r repository.EventRepository) EventService {
	return &eventService{r}
}

func (s *eventService) GetEvents() ([]dto.EventResponse, error) {
	events, err := s.repo.GetEvents()
	if err != nil {
		return nil, err
	}

	return dto.ToEventResponses(&events), nil
}

func (s *eventService) GetEventByID(id string) (*dto.EventResponse, error) {
	event, err := s.repo.GetEventByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("resource not found")
		}
		return nil, err
	}

	return dto.ToEventResponse(&event), nil
}

func (s *eventService) CreateEvent(id string, arg dto.CreateEventReq) (*dto.ResponseID, error) {
	date, err := time.Parse("2006-01-02", arg.Date)
	if err != nil {
		return nil, err
	}
	startTime, err := time.Parse("15:04", arg.StartTime)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse("15:04", arg.EndTime)
	if err != nil {
		return nil, err
	}

	startTime = time.Date(date.Year(), date.Month(), date.Day(), startTime.Hour(), startTime.Minute(), 0, 0, time.Local)
	endTime = time.Date(date.Year(), date.Month(), date.Day(), endTime.Hour(), endTime.Minute(), 0, 0, time.Local)

	input := sqlc.CreateEventParams{
		ID:             uuid.New().String(),
		OrganizationID: id,
		Title:          arg.Title,
		Description:    sql.NullString{String: arg.Description, Valid: arg.Description != ""},
		MaxRegistrant:  sql.NullInt32{Int32: arg.MaxRegistrant, Valid: true},
		Date:           date,
		StartTime:      sql.NullTime{Time: startTime, Valid: !startTime.IsZero()},
		EndTime:        sql.NullTime{Time: endTime, Valid: !endTime.IsZero()},
	}

	_, err = s.repo.CreateEvent(input)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return dto.NewResponseID(input.ID), nil
}

func (s *eventService) GetEventsByCategory(ids []string) ([]dto.EventResponse, error) {
	ctg_ids := strings.Join(ids, ",")

	var events []sqlc.Event
	data, err := s.repo.GetEventsByCategory(ctg_ids)
	if err != nil {
		return nil, err
	}
	events = append(events, data...)

	return dto.ToEventResponses(&events), nil
}

func (s *eventService) DeleteEvent(id string, org_id string) error {
	event, err := s.repo.GetEventByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("resource not found")
		}
		return err
	}

	if event.OrganizationID != org_id {
		return errors.New("access denied, event does not belong to user")
	}

	err = s.repo.DeleteEvent(id)
	return err
}

package service

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
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
	GetPricingsForEvent(string) ([]dto.PricingResponse, error)
	CreateEventPricing(dto.PricingCreateReq, string) (*dto.ResponseID, error)
	GetSpeakersForEvent(string) ([]dto.SpeakerResponse, error)
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

	eventResponses := []dto.EventResponse{}
	for _, event := range events {
		pricings, err := s.repo.GetPricingsForEvent(event.ID)
		if err != nil {
			return nil, err
		}
		speakers, err := s.repo.GetSpeakersForEvent(event.ID)
		if err != nil {
			return nil, err
		}
		eventResponse := dto.ToEventResponse(&event, &pricings, &speakers)
		eventResponses = append(eventResponses, *eventResponse)
	}

	return eventResponses, nil
}

func (s *eventService) GetEventByID(id string) (*dto.EventResponse, error) {
	event, err := s.repo.GetEventByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("event resource not found")
		}
		return nil, err
	}

	pricings, err := s.repo.GetPricingsForEvent(event.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("pricing resources not found")
		}
		return nil, err
	}

	speakers, err := s.repo.GetSpeakersForEvent(event.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("speakers resources not found")
		}
		return nil, err
	}

	return dto.ToEventResponse(&event, &pricings, &speakers), nil
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

func (s *eventService) GetPricingsForEvent(event_id string) ([]dto.PricingResponse, error) {
	pricings, err := s.repo.GetPricingsForEvent(event_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("resources not found")
		}
		return nil, err
	}
	return dto.ToPricingResponses(&pricings), nil
}

func (s *eventService) CreateEventPricing(arg dto.PricingCreateReq, event_id string) (*dto.ResponseID, error) {
	input := sqlc.CreateEventPricingParams{
		EventID:   event_id,
		EventType: sql.NullString{String: arg.EventType, Valid: true},
		Price:     sql.NullInt32{Int32: arg.Price, Valid: true},
	}

	ret, err := s.repo.CreateEventPricing(input)
	if err != nil {
		return nil, err
	}

	idRet, err := ret.LastInsertId()
	if err != nil {
		return nil, err
	}
	return dto.NewResponseID(strconv.FormatInt(idRet, 10)), nil
}

func (s *eventService) GetSpeakersForEvent(event_id string) ([]dto.SpeakerResponse, error) {
	speakers, err := s.repo.GetSpeakersForEvent(event_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("resources not found")
		}
		return nil, err
	}
	return dto.ToSpeakerResponses(&speakers), nil
}

func (s *eventService) CreateSpeaker(arg dto.SpeakerCreateReq, event_id string) (*dto.ResponseID, error) {
	input := sqlc.CreateSpeakerParams{
		ID:          uuid.New().String(),
		EventID:     sql.NullString{String: event_id, Valid: true},
		Name:        arg.Name,
		Title:       sql.NullString{String: arg.Name, Valid: arg.Title != ""},
		Description: sql.NullString{String: arg.Description, Valid: true},
	}

	_, err := s.repo.CreateSpeaker(input)
	if err != nil {
		return nil, err
	}

	return dto.NewResponseID(input.ID), nil
}

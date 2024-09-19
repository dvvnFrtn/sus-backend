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
	CreateEventAgenda(string, dto.CreateAgendaReq) (*dto.ResponseID, error)
	GetPricingsForEvent(string) ([]dto.PricingResponse, error)
	GetAgendasByEventID(string) ([]dto.AgendaResponse, error)
	CreateEventPricing(string, dto.PricingCreateReq) (*dto.ResponseID, error)
	GetSpeakersForAgenda(string) ([]dto.SpeakerResponse, error)
	CreateSpeaker(string, dto.SpeakerCreateReq) (*dto.ResponseID, error)
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
			return nil, errors.New("event resource not found")
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

	input := sqlc.CreateEventParams{
		ID:             uuid.New().String(),
		OrganizationID: id,
		Title:          arg.Title,
		Description:    sql.NullString{String: arg.Description, Valid: arg.Description != ""},
		MaxRegistrant:  sql.NullInt32{Int32: arg.MaxRegistrant, Valid: true},
		Date:           date,
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

func (s *eventService) CreateEventPricing(event_id string, arg dto.PricingCreateReq) (*dto.ResponseID, error) {
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

func (s *eventService) GetAgendasByEventID(event_id string) ([]dto.AgendaResponse, error) {
	agendas, err := s.repo.GetAgendasByEventID(event_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("resources not found")
		}
		return nil, err
	}
	return dto.ToAgendaResponses(&agendas), nil
}

func (s *eventService) CreateEventAgenda(event_id string, arg dto.CreateAgendaReq) (*dto.ResponseID, error) {
	startTime, err := time.Parse("2006-01-02 15:04", arg.StartTime)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse("2006-01-02 15:04", arg.EndTime)
	if err != nil {
		return nil, err
	}

	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(), startTime.Minute(), 0, 0, time.Local)
	endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), endTime.Hour(), endTime.Minute(), 0, 0, time.Local)

	input := sqlc.CreateEventAgendaParams{
		ID:          uuid.New().String(),
		EventID:     event_id,
		Title:       sql.NullString{String: arg.Title, Valid: true},
		Description: sql.NullString{String: arg.Description, Valid: true},
		StartTime:   sql.NullTime{Time: startTime, Valid: !startTime.IsZero()},
		EndTime:     sql.NullTime{Time: endTime, Valid: !endTime.IsZero()},
		Location:    sql.NullString{String: arg.Location, Valid: true},
	}

	_, err = s.repo.CreateEventAgenda(input)
	if err != nil {
		return nil, err
	}

	return dto.NewResponseID(input.ID), nil
}

func (s *eventService) GetSpeakersForAgenda(agenda_id string) ([]dto.SpeakerResponse, error) {
	speakers, err := s.repo.GetSpeakersForAgenda(agenda_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("resources not found")
		}
		return nil, err
	}
	return dto.ToSpeakerResponses(&speakers), nil
}

func (s *eventService) CreateSpeaker(agenda_id string, arg dto.SpeakerCreateReq) (*dto.ResponseID, error) {
	input := sqlc.CreateSpeakerParams{
		ID:          uuid.New().String(),
		AgendaID:    agenda_id,
		Name:        arg.Name,
		Title:       sql.NullString{String: arg.Name, Valid: true},
		Description: sql.NullString{String: arg.Description, Valid: true},
	}

	_, err := s.repo.CreateSpeaker(input)
	if err != nil {
		return nil, err
	}

	return dto.NewResponseID(input.ID), nil
}

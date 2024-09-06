package dto

import (
	"sus-backend/internal/db/sqlc"
	"time"
)

type EventResponse struct {
	ID             string
	OrganizationID string
	Title          string
	Img            string
	Description    string
	Registrant     int32
	MaxRegistrant  int32
	Date           time.Time
	StartTime      time.Time
	EndTime        time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func ToEventResponse(event *sqlc.Event) *EventResponse {
	return &EventResponse{
		ID:             event.ID,
		OrganizationID: event.OrganizationID,
		Title:          event.Title,
		Img:            event.Img.String,
		Description:    event.Description.String,
		Registrant:     event.Registrant.Int32,
		MaxRegistrant:  event.MaxRegistrant.Int32,
		Date:           event.Date,
		StartTime:      event.StartTime.Time,
		EndTime:        event.EndTime.Time,
		CreatedAt:      event.CreatedAt.Time,
		UpdatedAt:      event.UpdatedAt.Time,
	}
}

func ToEventResponses(events *[]sqlc.Event) []EventResponse {
	eventResponses := []EventResponse{}
	for _, event := range *events {
		eventResponse := EventResponse{
			ID:             event.ID,
			OrganizationID: event.OrganizationID,
			Title:          event.Title,
			Img:            event.Img.String,
			Description:    event.Description.String,
			Registrant:     event.Registrant.Int32,
			MaxRegistrant:  event.MaxRegistrant.Int32,
			Date:           event.Date,
			StartTime:      event.StartTime.Time,
			EndTime:        event.EndTime.Time,
			CreatedAt:      event.CreatedAt.Time,
			UpdatedAt:      event.UpdatedAt.Time,
		}

		eventResponses = append(eventResponses, eventResponse)
	}
	return eventResponses
}

type CreateEventReq struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	MaxRegistrant int32  `json:"max_registrant"`
	Date          string `json:"date"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
}

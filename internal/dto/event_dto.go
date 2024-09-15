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

func ToEventResponse(event *sqlc.Event, pricings *[]sqlc.EventPricing, speakers *[]sqlc.Speaker) *EventResponse {
	return &EventResponse{
		ID:             event.ID,
		OrganizationID: event.OrganizationID,
		Title:          event.Title,
		Img:            event.Img.String,
		Description:    event.Description.String,
		Registrant:     event.Registrant.Int32,
		MaxRegistrant:  event.MaxRegistrant.Int32,
		Date:           event.Date,
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

type PricingResponse struct {
	ID        int32
	EventID   string
	EventType string
	Price     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToPricingResponse(pricing sqlc.EventPricing) *PricingResponse {
	return &PricingResponse{
		ID:        int32(pricing.ID),
		EventID:   pricing.EventID,
		EventType: pricing.EventType.String,
		Price:     pricing.Price.Int32,
		CreatedAt: pricing.CreatedAt.Time,
		UpdatedAt: pricing.UpdatedAt.Time,
	}
}

func ToPricingResponses(pricings *[]sqlc.EventPricing) []PricingResponse {
	pricingResponses := []PricingResponse{}
	for _, pricing := range *pricings {
		pricingResponse := PricingResponse{
			ID:        int32(pricing.ID),
			EventID:   pricing.EventID,
			EventType: pricing.EventType.String,
			Price:     pricing.Price.Int32,
			CreatedAt: pricing.CreatedAt.Time,
			UpdatedAt: pricing.UpdatedAt.Time,
		}

		pricingResponses = append(pricingResponses, pricingResponse)
	}
	return pricingResponses
}

type PricingCreateReq struct {
	EventType string `json:"event_type"`
	Price     int32  `json:"price"`
}

type SpeakerResponse struct {
	ID          string
	AgendaID    string
	Name        string
	Title       string
	Img         string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToSpeakerResponse(speaker sqlc.Speaker) *SpeakerResponse {
	return &SpeakerResponse{
		ID:          speaker.ID,
		AgendaID:    speaker.AgendaID,
		Name:        speaker.Name,
		Title:       speaker.Title.String,
		Img:         speaker.Img.String,
		Description: speaker.Description.String,
		CreatedAt:   speaker.CreatedAt.Time,
		UpdatedAt:   speaker.UpdatedAt.Time,
	}
}

func ToSpeakerResponses(speakers *[]sqlc.Speaker) []SpeakerResponse {
	speakerResponses := []SpeakerResponse{}
	for _, speaker := range *speakers {
		speakerResponse := SpeakerResponse{
			ID:          speaker.ID,
			AgendaID:    speaker.AgendaID,
			Name:        speaker.Name,
			Title:       speaker.Title.String,
			Img:         speaker.Img.String,
			Description: speaker.Description.String,
			CreatedAt:   speaker.CreatedAt.Time,
			UpdatedAt:   speaker.UpdatedAt.Time,
		}

		speakerResponses = append(speakerResponses, speakerResponse)
	}
	return speakerResponses
}

type SpeakerCreateReq struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

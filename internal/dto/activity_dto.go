package dto

import (
	"sus-backend/internal/db/sqlc"
	"time"
)

type ActivityCreateReq struct {
	Title     string `json:"title"`
	Note      string `json:"note"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type ActivityResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Note      string    `json:"note"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToActivityResponse(actv *sqlc.Activity) *ActivityResponse {
	return &ActivityResponse{
		ID:        actv.ID,
		Title:     actv.Title.String,
		Note:      actv.Note,
		StartTime: actv.StartTime.Time,
		EndTime:   actv.EndTime.Time,
		CreatedAt: actv.CreatedAt.Time,
		UpdatedAt: actv.UpdatedAt.Time,
	}
}

func ToActivityResponses(actvs *[]sqlc.Activity) []ActivityResponse {
	activityResponses := []ActivityResponse{}
	for _, actv := range *actvs {
		actvResponse := ActivityResponse{
			ID:        actv.ID,
			Title:     actv.Title.String,
			Note:      actv.Note,
			StartTime: actv.StartTime.Time,
			EndTime:   actv.EndTime.Time,
			CreatedAt: actv.CreatedAt.Time,
			UpdatedAt: actv.UpdatedAt.Time,
		}
		activityResponses = append(activityResponses, actvResponse)
	}
	return activityResponses
}

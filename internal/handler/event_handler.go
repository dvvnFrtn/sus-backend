package handler

import (
	"sus-backend/internal/dto"
	"sus-backend/internal/service"
	"sus-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	serv service.EventService
}

func NewEventHandler(s service.EventService) *EventHandler {
	return &EventHandler{s}
}

func (h *EventHandler) GetEvents(c *gin.Context) {
	data, err := h.serv.GetEvents()
	if err != nil {
		response.FailOrError(c, 500, "Failed getting events", err)
		return
	}
	response.Success(c, 200, "Success getting events", data)
}

func (h *EventHandler) GetEventByID(c *gin.Context) {
	idReq := c.Param("id")
	data, err := h.serv.GetEventByID(idReq)
	if err != nil {
		response.FailOrError(c, 500, "Failed getting event", err)
		return
	}
	response.Success(c, 200, "Success getting events", data)
}

func (h *EventHandler) AddEvent(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)

	var req dto.CreateEventReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailOrError(c, 400, "Bad request", err)
		return
	}

	resp, err := h.serv.CreateEvent(claims.ID, req)
	if err != nil {
		response.FailOrError(c, 500, "Failed creating event", err)
		return
	}
	response.Success(c, 201, "Success creating event", resp)
}

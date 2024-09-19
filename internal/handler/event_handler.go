package handler

import (
	"net/http"
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
	var req dto.RequestIDs
	_ = c.ShouldBindJSON(&req)

	var result []dto.EventResponse
	if req.IDs != nil {
		data, err := h.serv.GetEventsByCategory(req.IDs)
		if err != nil {
			response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
			return
		}
		result = data
	} else {
		data, err := h.serv.GetEvents()
		if err != nil {
			response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
			return
		}
		result = data
	}

	response.Success(c, http.StatusOK, "Resources Successfully Retrieved", result)
}

func (h *EventHandler) GetEventByID(c *gin.Context) {
	idReq := c.Param("id")
	data, err := h.serv.GetEventByID(idReq)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}
	response.Success(c, http.StatusOK, "Resource Successfully Retrieved", data)
}

func (h *EventHandler) AddEvent(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)

	var req dto.CreateEventReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request", err)
		return
	}

	resp, err := h.serv.CreateEvent(claims.ID, req)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}
	response.Success(c, http.StatusCreated, "Resource Created Successfully", resp)
}

func (h *EventHandler) DeleteEvent(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)

	idReq := c.Param("id")
	err := h.serv.DeleteEvent(idReq, claims.ID)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}
	response.Success(c, http.StatusOK, "Resource Deleted Successfully", nil)
}

func (h *EventHandler) GetPricingsByEventID(c *gin.Context) {
	idReq := c.Param("id")
	data, err := h.serv.GetPricingsForEvent(idReq)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Speakers retrieved successfully", data)
}

func (h *EventHandler) GetAgendasByEventID(c *gin.Context) {
	idReq := c.Param("id")
	data, err := h.serv.GetAgendasByEventID(idReq)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Agendas retrieved successfully", data)
}

func (h *EventHandler) CreateAgenda(c *gin.Context) {
	idReq := c.Param("id")
	var req dto.CreateAgendaReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request", err)
		return
	}

	resultID, err := h.serv.CreateEventAgenda(idReq, req)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(c, http.StatusCreated, "Agenda created successfully", resultID)
}

func (h *EventHandler) GetSpeakersByAgendaID(c *gin.Context) {
	idReq := c.Param("agendaid")
	data, err := h.serv.GetSpeakersForAgenda(idReq)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Speakers retrieved successfully", data)
}

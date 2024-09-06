package handler

import (
	"log"
	"net/http"
	"sus-backend/internal/dto"
	"sus-backend/internal/service"
	"sus-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type ActivityHandler struct {
	serv service.ActivityService
}

func NewActivityHandler(s service.ActivityService) *ActivityHandler {
	return &ActivityHandler{s}
}

func (h *ActivityHandler) GetActivityByID(c *gin.Context) {
	idReq := c.Param("id")
	data, err := h.serv.GetActivityByID(idReq)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}
	response.Success(c, http.StatusOK, "Resource Successfully Retrieved", data)
}

func (h *ActivityHandler) GetActivitiesByOrganizationID(c *gin.Context) {
	idReq := c.Param("id")
	data, err := h.serv.GetActivitiesByOrganizationID(idReq)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}
	response.Success(c, http.StatusOK, "Resource Successfully Retrieved", data)
}

func (h *ActivityHandler) CreateActivity(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)
	log.Println(claims)

	var req dto.ActivityCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailOrError(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	result, err := h.serv.CreateActivity(req, claims.OrganizationID)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}
	response.Success(c, http.StatusCreated, "Resource Created Successfully", result)
}

func (h *ActivityHandler) DeleteActivity(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)

	idReq := c.Param("id")
	err := h.serv.DeleteActivity(idReq, claims.OrganizationID)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}
	response.Success(c, http.StatusOK, "Resource Successfully Deleted", nil)
}

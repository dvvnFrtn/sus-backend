package handler

import (
	"fmt"
	"net/http"
	"sus-backend/internal/dto"
	"sus-backend/internal/service"
	"sus-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	serv service.OrganizationService
}

func NewOrganizationHandler(serv service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{serv}
}

func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var organizationReq dto.OrganizationCreateRequest
	err := c.ShouldBindJSON(&organizationReq)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Error Bad Request", nil)
		return
	}

	res, err := h.serv.CreateOrganization(organizationReq)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Oops Something Went Wrong", err)
		return
	}

	response.Success(c, http.StatusCreated, "Resource Created Successfully", res)
}

func (h *OrganizationHandler) FindOrganizationById(c *gin.Context) {
	idReq := c.Param("id")
	res, err := h.serv.FindOrganizationById(idReq)
	if err != nil {
		fmt.Println(err.Error())
		response.FailOrError(c, http.StatusNotFound, "Resource Not Found", nil)
		return
	}

	response.Success(c, http.StatusFound, "Resource Successfully Retrieved", res)
}

func (h *OrganizationHandler) ListAllOrganizations(c *gin.Context) {
	res, err := h.serv.ListAllOrganizations()
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Oops Something Went Wrong", err)
		return
	}

	response.Success(c, http.StatusOK, "Resources Successfully Retrievied", res)
}

func (h *OrganizationHandler) UpdateOrganizations(c *gin.Context) {
	idReq := c.Param("id")
	var organizationReq dto.OrganizationUpdateRequest
	err := c.ShouldBindJSON(&organizationReq)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Error Bad Request", err)
		return
	}

	res, err := h.serv.UpdateOrganization(idReq, organizationReq)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Resource Not Found", err)
		return
	}

	response.Success(c, http.StatusOK, "Resource Updated Successfully", res)
}

func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
	idReq := c.Param("id")
	err := h.serv.DeleteOrganization(idReq)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Resource Not Found", err)
		return
	}

	response.Success(c, http.StatusNoContent, "Resource Deleted Successfully", nil)
}

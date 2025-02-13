package handler

import (
	"net/http"
	"sus-backend/internal/dto"
	"sus-backend/internal/service"
	_response "sus-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	serv service.OrganizationService
}

func NewOrganizationHandler(serv service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{serv}
}

func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)

	request := new(dto.OrganizationCreateRequest)
	if err := c.ShouldBindJSON(&request); err != nil {
		_response.FailOrError(c, http.StatusBadRequest, "invalid_request", err)
		return
	}

	response, err := h.serv.CreateOrganization(claims.ID, *request)
	if err != nil {
		_response.FailOrError(c, 500, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusCreated, "Organization Created Successfully", response)
}

func (h *OrganizationHandler) GetOrganizationById(c *gin.Context) {
	organizationID := c.Param("id")

	response, err := h.serv.GetOrganizationById(organizationID)
	if err != nil {
		_response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusFound, "Organization Retrieved Successfully", response)
}

func (h *OrganizationHandler) GetAllOrganizations(c *gin.Context) {
	response, err := h.serv.GetAllOrganizations()
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Organizations Retrieved Successfully", response)
}

func (h *OrganizationHandler) UpdateOrganizations(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)
	organizationID := c.Param("id")

	request := new(dto.OrganizationUpdateRequest)
	if err := c.ShouldBindJSON(&request); err != nil {
		_response.FailOrError(c, http.StatusBadRequest, "invalid_request", err)
		return
	}

	response, err := h.serv.UpdateOrganization(claims.ID, organizationID, *request)
	if err != nil {
		_response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Organization Updated Successfully", response)
}

func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)
	organizationID := c.Param("id")

	if err := h.serv.DeleteOrganization(claims.ID, organizationID); err != nil {
		_response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusNoContent, "Organization Deleted Successfully", nil)
}

func (h *OrganizationHandler) Follow(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)
	organizationID := c.Param("id")

	if err := h.serv.Follow(claims.ID, organizationID); err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Organization Followed Successfully", nil)
}

func (h *OrganizationHandler) Unfollow(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)
	organizationID := c.Param("id")

	if err := h.serv.Unfollow(claims.ID, organizationID); err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Organization Unfollowed Successfully", nil)
}

func (h *OrganizationHandler) GetFollowers(c *gin.Context) {
	organizationID := c.Param("id")

	response, err := h.serv.GetFollowers(organizationID)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Followers Retrieved Successfully", response)
}

func (h *OrganizationHandler) GetCategories(c *gin.Context) {
	data, err := h.serv.GetCategories()
	if err != nil {
		_response.FailOrError(c, 500, "Failed getting categories", err)
		return
	}
	_response.Success(c, 200, "Categories Retrieved Successfully", data)
}

func (h *OrganizationHandler) GetFollowedOrganizations(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)

	response, err := h.serv.GetFollowedOrganizations(claims.ID)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Organizations Retrieved Successfully", response)
}

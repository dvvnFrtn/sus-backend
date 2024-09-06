package handler

import (
	"net/http"
	"sus-backend/internal/dto"
	"sus-backend/internal/service"
	"sus-backend/pkg/response"
	_response "sus-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	serv service.PostService
}

func NewPostHandler(serv service.PostService) *PostHandler {
	return &PostHandler{serv}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	// TODO: implement using organizationID from claims
	organizationID := c.Query("org_id")

	request := new(dto.PostCreateRequest)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		_response.FailOrError(c, http.StatusBadRequest, "invalid_request", err)
		return
	}

	response, err := h.serv.CreatePost(organizationID, *request)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusCreated, "Resource Created Successfully", response)
}

func (h *PostHandler) FindPostById(c *gin.Context) {
	idReq := c.Param("id")
	res, err := h.serv.FindPostById(idReq)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}

	response.Success(c, http.StatusFound, "Resource Successfully Retrieved", res)
}

func (h *PostHandler) FindPostsByOrganization(c *gin.Context) {
	idReq := c.Param("id")
	res, err := h.serv.FindPostByOrganization(idReq)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Resources Successfully Retrievied", res)
}

func (h *PostHandler) ListAllPosts(c *gin.Context) {
	res, err := h.serv.ListAllPosts()
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Resources Successfully Retrievied", res)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	postID := c.Param("id")
	organizationID := c.Query("org_id")

	if err := h.serv.DeletePost(organizationID, postID); err != nil {
		_response.FailOrError(c, http.StatusNotFound, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusNoContent, "Resource Deleted Successfully", nil)
}

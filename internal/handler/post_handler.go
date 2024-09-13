package handler

import (
	"fmt"
	"net/http"
	"sus-backend/internal/dto"
	"sus-backend/internal/service"
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
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)

	request := new(dto.PostCreateRequest)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		_response.FailOrError(c, http.StatusBadRequest, "invalid_request", err)
		return
	}

	response, err := h.serv.CreatePost(claims.OrganizationID, *request)
	if err != nil {
		fmt.Println(err)
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusCreated, "Resource Created Successfully", response)
}

func (h *PostHandler) GetPostById(c *gin.Context) {
	postID := c.Param("id")

	response, err := h.serv.GetPostById(postID)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusFound, "Resource Successfully Retrieved", response)
}

func (h *PostHandler) GetPostsByOrganization(c *gin.Context) {
	organizationID := c.Param("id")

	response, err := h.serv.GetPostsByOrganization(organizationID)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Resources Successfully Retrievied", response)
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	response, err := h.serv.GetAllPosts()
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Resources Successfully Retrievied", response)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)
	postID := c.Param("id")

	if err := h.serv.DeletePost(claims.OrganizationID, postID); err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusNoContent, "Resource Deleted Successfully", nil)
}

func (h *PostHandler) LikedPost(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)
	postID := c.Param("id")

	if err := h.serv.LikedPost(claims.ID, postID); err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Post Liked Successfully", nil)
}

func (h *PostHandler) UnlikedPost(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)
	postID := c.Param("id")

	if err := h.serv.UnlikedPost(claims.ID, postID); err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Post Unliked Successfully", nil)
}

func (h *PostHandler) GetPostLikes(c *gin.Context) {
	postID := c.Param("id")

	response, err := h.serv.GetPostLikes(postID)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Resource Retrievied Successfully", response)
}

func (h *PostHandler) CommentPost(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)
	postID := c.Param("id")

	request := new(dto.CommentPostRequest)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		_response.FailOrError(c, http.StatusBadRequest, "invalid_request", err)
		return
	}

	response, err := h.serv.CommentPost(claims.ID, postID, *request)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusCreated, "Resource Created Successfully", response)
}

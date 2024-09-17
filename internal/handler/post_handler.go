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

	_response.Success(c, http.StatusCreated, "Post Created Successfully", response)
}

func (h *PostHandler) GetPostById(c *gin.Context) {
	postID := c.Param("id")

	response, err := h.serv.GetPostById(postID)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusFound, "Post Retrieved Successfully", response)
}

func (h *PostHandler) GetPostsByOrganization(c *gin.Context) {
	organizationID := c.Param("id")

	response, err := h.serv.GetPostsByOrganization(organizationID)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Posts Retrieved Successfully", response)
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)

	response, err := h.serv.GetAllPosts(claims.ID)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Posts Retrieved Successfully", response)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)
	postID := c.Param("id")

	if err := h.serv.DeletePost(claims.OrganizationID, postID); err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusNoContent, "Post Deleted Successfully", nil)
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

	_response.Success(c, http.StatusOK, "Likes Retrieved Successfully", response)
}

func (h *PostHandler) CommentPost(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)

	request := new(dto.CommentPostRequest)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		_response.FailOrError(c, http.StatusBadRequest, "invalid_request", err)
		return
	}

	response, err := h.serv.CommentPost(claims.ID, *request)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusCreated, "Comment Created Successfully", response)
}

func (h *PostHandler) GetPostComments(c *gin.Context) {
	postID := c.Param("id")

	response, err := h.serv.GetPostComments(postID)
	if err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusOK, "Comments Retrieved Successfully", response)
}

func (h *PostHandler) DeleteComment(c *gin.Context) {
	auth, _ := c.Get("user")
	claims := auth.(*dto.UserClaims)
	commentID := c.Param("id")

	if err := h.serv.DeleteComment(claims.ID, commentID); err != nil {
		_response.FailOrError(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	_response.Success(c, http.StatusNoContent, "Comment Deleted Successfully", nil)
}

package handler

import (
	"database/sql"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/repository"
	"sus-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func route(r *gin.Engine, uh *UserHandler, oh *OrganizationHandler, ph *PostHandler) {
	r.GET("/auth/google/login-w-google", uh.LoginWithGoogle)
	r.GET("/auth/google/callback", uh.GetGoogleDetails)
	r.POST("/register", uh.RegisterUser)
	r.GET("/account-confirm", uh.CreateConfirmedUser)
	r.POST("/login", uh.Login)

	r.POST("/organizations", oh.CreateOrganization)
	r.GET("/organizations/:id/posts", ph.FindPostsByOrganization)
	r.GET("/organizations/:id", oh.FindOrganizationById)
	r.GET("/organizations", oh.ListAllOrganizations)
	r.PUT("/organizations/:id", oh.UpdateOrganizations)
	r.DELETE("/organizations/:id", oh.DeleteOrganization)

	r.POST("/posts", ph.CreatePost)
	r.GET("/posts", ph.ListAllPosts)
	r.GET("/posts/:id", ph.FindPostById)
	r.DELETE("/posts/:id", ph.DeletePost)
}

func InitHandler(db *sql.DB) (*UserHandler, *OrganizationHandler, *PostHandler) {
	queries := sqlc.New(db)

	userRepo := repository.NewUserRepository(queries)
	userServ := service.NewUserService(userRepo)
	userHand := NewUserHandler(userServ)

	organizationRepo := repository.NewOrganizationRepository(queries)
	organizationServ := service.NewOrganizationService(organizationRepo)
	organizationHand := NewOrganizationHandler(organizationServ)

	postRepo := repository.NewPostRepository(queries)
	postServ := service.NewPostService(postRepo, organizationServ)
	postHand := NewPostHandler(postServ)

	return userHand, organizationHand, postHand
}

func StartEngine(r *gin.Engine, db *sql.DB) {
	uh, oh, ph := InitHandler(db)
	route(r, uh, oh, ph)
}

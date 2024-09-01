package handler

import (
	"database/sql"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/repository"
	"sus-backend/internal/service"
	"sus-backend/middleware"

	"github.com/gin-gonic/gin"
)

func route(r *gin.Engine, uh *UserHandler, oh *OrganizationHandler, eh *EventHandler) {
	r.GET("/auth/google/login-w-google", uh.LoginWithGoogle)
	r.GET("/auth/google/callback", uh.GetGoogleDetails)
	r.POST("/register", uh.RegisterUser)
	r.GET("/account-confirm", uh.CreateConfirmedUser)
	r.POST("/login", uh.Login)
	r.GET("/users/:id", uh.FindUserByID)
	r.PUT("/users", middleware.ValidateToken("user"), uh.UpdateUser)
	r.POST("/user-categories", middleware.ValidateToken("user"), uh.AddUserCategory)

	r.POST("/organizations", oh.CreateOrganization)
	r.GET("/organizations/:id", oh.FindOrganizationById)
	r.GET("/organizations", oh.ListAllOrganizations)
	r.PUT("/organizations/:id", oh.UpdateOrganizations)
	r.DELETE("/organizations/:id", oh.DeleteOrganization)

	r.GET("/events", eh.GetEvents)
	r.GET("/events/:id", eh.GetEventByID)
	r.POST("/events", middleware.ValidateToken("organization"), eh.AddEvent)
}

func InitHandler(db *sql.DB) (*UserHandler, *OrganizationHandler, *EventHandler) {
	queries := sqlc.New(db)

	userRepo := repository.NewUserRepository(queries)
	userServ := service.NewUserService(userRepo)
	userHand := NewUserHandler(userServ)

	organizationRepo := repository.NewOrganizationRepository(queries)
	organizationServ := service.NewOrganizationService(organizationRepo)
	organizationHand := NewOrganizationHandler(organizationServ)

	eventRepo := repository.NewEventRepository(queries)
	eventServ := service.NewEventService(eventRepo)
	eventHand := NewEventHandler(eventServ)

	return userHand, organizationHand, eventHand
}

func StartEngine(r *gin.Engine, db *sql.DB) {
	uh, oh, eh := InitHandler(db)
	route(r, uh, oh, eh)
}

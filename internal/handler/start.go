package handler

import (
	"database/sql"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/repository"
	"sus-backend/internal/service"
	"sus-backend/middleware"

	"github.com/gin-gonic/gin"
)

func route(r *gin.Engine, uh *UserHandler, oh *OrganizationHandler, ph *PostHandler, eh *EventHandler, ah *ActivityHandler) {
	r.GET("/auth/google/login-w-google", uh.LoginWithGoogle)
	r.GET("/auth/google/callback", uh.GetGoogleDetails)
	r.POST("/register", uh.RegisterUser)
	r.GET("/account-confirm", uh.CreateConfirmedUser)
	r.POST("/login", uh.Login)
	r.GET("/users/:id", uh.FindUserByID)
	r.PUT("/users", middleware.ValidateToken("user"), uh.UpdateUser)
	r.POST("/user-categories", middleware.ValidateToken("user"), uh.AddUserCategory)
	r.POST("/login-organizers", uh.LoginForOrganizer)

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

	r.GET("/events", eh.GetEvents)
	r.GET("/events/:id", eh.GetEventByID)
	r.POST("/events", middleware.ValidateToken("organization"), eh.AddEvent)
	r.DELETE("events/:id", middleware.ValidateToken("organization"), eh.DeleteEvent)

	r.GET("/activities/:id", ah.GetActivityByID)
	r.GET("/organizations/:id/activities", ah.GetActivitiesByOrganizationID)
	r.POST("/activities", middleware.ValidateToken("organization"), ah.CreateActivity)
	r.DELETE("/activities/:id", middleware.ValidateToken("organization"), ah.DeleteActivity)
}

func InitHandler(db *sql.DB) (*UserHandler, *OrganizationHandler, *PostHandler, *EventHandler, *ActivityHandler) {
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

	eventRepo := repository.NewEventRepository(queries)
	eventServ := service.NewEventService(eventRepo)
	eventHand := NewEventHandler(eventServ)

	activityRepo := repository.NewActivityRepository(queries)
	activityServ := service.NewActivityService(activityRepo)
	activityHand := NewActivityHandler(activityServ)

	return userHand, organizationHand, postHand, eventHand, activityHand
}

func StartEngine(r *gin.Engine, db *sql.DB) {
	uh, oh, ph, eh, ah := InitHandler(db)
	route(r, uh, oh, ph, eh, ah)
}

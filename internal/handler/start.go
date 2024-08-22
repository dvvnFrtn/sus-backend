package handler

import (
	"database/sql"
	"sus-backend/internal/db/sqlc"
	"sus-backend/internal/repository"
	"sus-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func route(r *gin.Engine, uh *UserHandler) {
	r.GET("/auth/google/login-w-google", uh.LoginWithGoogle)
	r.GET("/auth/google/callback", uh.GetGoogleDetails)
	r.POST("/register", uh.RegisterUser)
	r.GET("/account-confirm", uh.CreateConfirmedUser)
	r.POST("/login", uh.Login)
}

func InitHandler(db *sql.DB) *UserHandler {
	queries := sqlc.New(db)

	userRepo := repository.NewUserRepository(queries)
	userServ := service.NewUserService(userRepo)
	userHand := NewUserHandler(userServ)

	return userHand
}

func StartEngine(r *gin.Engine, db *sql.DB) {
	uh := InitHandler(db)
	route(r, uh)
}

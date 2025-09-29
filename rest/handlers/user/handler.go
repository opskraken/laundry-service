package user

import (
	"github.com/enghasib/laundry_service/config"
	userRepo "github.com/enghasib/laundry_service/repo/user"
	middleware "github.com/enghasib/laundry_service/rest/middlewares"
)

type UserHandler struct {
	cnf        *config.Config
	middleware middleware.Middlewares
	usrRepo    userRepo.UserRepo
}

func NewUserHandler(
	cnf *config.Config,
	middleware middleware.Middlewares,
	usrRepo userRepo.UserRepo,

) *UserHandler {
	return &UserHandler{
		cnf:        cnf,
		middleware: middleware,
		usrRepo:    usrRepo,
	}
}

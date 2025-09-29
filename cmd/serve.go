package cmd

import (
	"fmt"
	"os"

	"github.com/enghasib/laundry_service/config"
	"github.com/enghasib/laundry_service/db"
	userRepo "github.com/enghasib/laundry_service/repo/user"
	"github.com/enghasib/laundry_service/rest"
	"github.com/enghasib/laundry_service/rest/handlers/user"
	middleware "github.com/enghasib/laundry_service/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	db, err := db.NewConnection(cnf)
	if err != nil {
		fmt.Println("DB connection Error, Error:", err)
		os.Exit(1)
	}

	usrRepo := userRepo.NewUserRepo(db)

	middlewares := middleware.NewMiddlewares(cnf)

	userHandler := user.NewUserHandler(cnf, *middlewares, usrRepo)

	server := rest.NewServer(*cnf, *userHandler)
	server.Start()

}

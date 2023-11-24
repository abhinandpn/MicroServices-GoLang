// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/api"
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/api/handler"
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/config"
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/database"
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/repositories"
	"github.com/abhinandpn/MicroServices-GoLang/user-services/pkg/usecase"
)

// Injectors from wire.go:

func InitializeApi(cfg config.Config) (*http.ServeHTTP, error) {
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	serveHTTP := http.NewServeHttp(userHandler)
	return serveHTTP, nil
}

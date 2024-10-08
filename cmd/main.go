package main

import (
	"app-server/cmd/inject"
	"app-server/internal/infrastructure/config"
	"log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	//server, err := InitializeServer(config)
	//db, err := database.Connect(config)
	//if err != nil {
	//log.Fatalf("Failed to connect to database: %v", err)
	//}
	//userRepository := postgres.NewUserRepository(db)
	//userServiceInterface := user.NewService(userRepository)
	//userHandler := v1.NewUserHandler(userServiceInterface)
	//
	//// Account-related dependencies
	//authService := auth.NewAuthService(*config)
	//accountRepository := repository.NewGenericBaseRepository[entity.UserRole](db)
	//accountServiceInterface := account.NewAccountService(userRepository, accountRepository, authService)
	//accountHandler := v1.NewAccountHandler(accountServiceInterface)

	//server := server.NewHTTPServer(config, userHandler, accountHandler, authService)
	server, err := inject.InitializeServer(config)
	if err != nil {
		log.Fatalf("Failed to initialize API: %v", err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

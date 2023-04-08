package main

import (

	"log"
	"net/http"
	"bityagi/logger"
	"bityagi/service"
	"bityagi/handlers"

	openapi "build/code/spec/src"
)

func main() {
	// loadEnv()
	// loadDatabase()
	serveApplication()

}

// func loadEnv() {
// 	err := godotenv.Load(".env.local")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

// func loadDatabase() {
// 	database.Connect()
// 	database.Database.AutoMigrate(&model.User{})
// 	database.Database.AutoMigrate(&model.Entry{})
// }

func serveApplication() {
	logger.InitialzieLogger()

	PersonApiService := service.NewMyPersonApiService()

	PersonApiController := openapi.NewPersonApiController(PersonApiService)

	router := openapi.NewRouter(PersonApiController)

	apiDocHandler := handlers.NewApiDocHandler()
	router.Handle("/*", apiDocHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
	logger.Logger.Info("Starting Application in port 8080")


}
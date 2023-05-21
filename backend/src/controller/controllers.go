package controller

import (
	"bityagi/service"
	openapi "build/code/spec/src"
	"github.com/go-chi/chi/v5"
)

func SetupServicesAndControllers() (chi.Router, error) {

	// Create services
	personApiService := service.NewMyPersonApiService()
	newMyPetApiService := service.NewMyPetApiService()

	// Set up API controllers
	personApiController := openapi.NewPersonAPIController(personApiService)
	petsApiController := openapi.NewPetsAPIController(newMyPetApiService)

	// Set up router
	router := openapi.NewRouter(personApiController, petsApiController)

	return router, nil
}

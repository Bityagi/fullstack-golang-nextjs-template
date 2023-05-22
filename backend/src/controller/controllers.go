package controller

import (
	"bityagi/handlers"
	openapi "build/code/spec/src"
	"github.com/go-chi/chi/v5"
)

func SetupServicesAndControllers() (chi.Router, error) {

	// Create services
	newMyPetApiService := handlers.NewMyPetApiService()

	// Set up API controllers
	petsApiController := openapi.NewPetsAPIController(newMyPetApiService)

	// Set up router, if you have another controller you add it here
	// Example - router := openapi.NewRouter(personApiController, petsApiController, newController)
	router := openapi.NewRouter(petsApiController)

	return router, nil
}

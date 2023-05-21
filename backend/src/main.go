package main

import (
	"bityagi/controller"
	"bityagi/handlers"
	"bityagi/logger"
	"log"
	"net/http"
)

func main() {
	serveApplication()
}

func serveApplication() {
	logger.InitialzieLogger()

	//All controllers are initialized with their services at ./controller/controller.go
	router, err := controller.SetupServicesAndControllers()

	//Health check handler, this need to match in your deployment files, the cluster will fail if it can't get a health check
	router.HandleFunc("/health", handlers.HealthCheckHandler)

	//Serves all static resources found at ../backend/spec, serving swagger.html file at root
	apiDocHandler := handlers.NewApiDocHandler()
	router.Handle("/*", apiDocHandler)

	logger.Logger.Info("Starting Application in port 8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

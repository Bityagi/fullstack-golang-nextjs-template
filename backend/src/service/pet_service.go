package service

import (
	"bityagi/logger"
	openapi "build/code/spec/src"
	"go.uber.org/zap"
	"strconv"
)

func GetPetByID(petID int64) openapi.Pet {
	// Perform the necessary operations to fetch the pet from the database or any other data source
	// In this example, a dummy pet is returned
	logger.Logger.Info("Person value is called", zap.String("pet", strconv.Itoa(int(petID))))
	return openapi.Pet{Id: petID, Name: "Dog", Tag: "hate"}
}

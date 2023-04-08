package service

import (
	openapi "build/code/spec/src"
	"bityagi/logger"
	"context"
	"go.uber.org/zap"
	"net/http"
)

type MyPersonApiService struct{}

func NewMyPersonApiService() openapi.PersonApiServicer {
	return &MyPersonApiService{}
}

func (s *MyPersonApiService) ShowPerson(ctx context.Context) (openapi.ImplResponse, error) {
	logger.InitialzieLogger()

	person := openapi.Person{Id: 23, Name: "Dog", Age: 23}

	logger.Logger.Info("Person value is called", zap.String("person", person.Name))

	return openapi.Response(http.StatusOK, person), nil
}



# Go OpenApi Template

## Table of Contents
- **[Overview](#overview)**
- **[Prerequisites](#prerequisites)**
- **[Build Controller & Models](#build-controller--models)**
- **[Build the Project](#build-the-project)**
- **[Run the Project](#run-the-project)**
- **[Writing OpenAPI Specs](#writing-openapi-specs)**
  - [Defining Endpoints](#defining-endpoints)
  - [Detailing Endpoint Methods](#detailing-endpoint-methods)
  - [Defining Schemas](#defining-schemas)
  - [Defining Responses](#defining-responses)
  - [Defining Parameters](#defining-parameters)
- **[Implementing the Services](#implementing-the-services)**
  - [Adding Business Logic](#adding-business-logic)
  - [Connecting the Service to the Controller](#connecting-the-service-to-the-controller)

## Overview
This is a Golang project that takes on an OpenAPI Spec first approach. This means the developers starts development by planning out the desired APIs with product or their product alter ego.
Once the APIs are defined, the developer can then generate the controller and models from the OpenAPI spec. The developer can then start implementing the business logic for the APIs.

## Prerequisites
- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)

## Build Controller & Models
```shell
make generateCode
```

## Build the Project
Go to the directory [/backend/src](/backend/src) and run
```shell
go build
```

## Run the Project
Go to the directory [/backend/src](/backend/src) and run
```shell
go run main.go
```

## URLs
**Swagger UI:** http://localhost:8080/swagger.html

**Redocly UI:** http://localhost:8080/redocly.html

**Health Check:** http://localhost:8080/health

## Docker Image
Running the following command will build the docker image and run it locally.
```shell
docker build -t bityagi .
```
TODO: Need to fix the docker image's swagger UI and redocly UI. The image can't serve the static files they aren't copying over properly. 

Running the following command will run the docker image locally.
```shell
docker run -p 8080:8080 bityagi
```

## Writing OpenAPI Specs
Write the specs in the file [backend/spec/openapi/openapi.yaml](./spec/openapi/openapi.yaml). The spec is written in the [OpenAPI 3.0](https://swagger.io/specification/) format and references multiple files in the [openapi](./spec/openapi) folder.

The open spec folder is organized as follows:
- [openapi-v1.0.yaml](./spec/openapi/openapi-v1.0.yaml): **(GENERATED - DO NOT EDIT)** The file that is generated from the [openapi.yaml](./spec/openapi/openapi.yaml) file. This file is used by the [OpenAPI UI](http://localhost:8080/docs) to display the API documentation.

- [openapi.yaml](./spec/openapi/openapi.yaml): **(EDIT THIS)** The file that references all the other files.
- [resources](./spec/openapi/resources): Contains all the endpoints defined in the spec. Each endpoint is defined in a separate file.
- [schemas](./spec/openapi/schemas): Contains all the schemas used in the spec.
- [responses](./spec/openapi/responses): Contains all the responses used in the spec.
- [parameters](./spec/openapi/parameters): Contains all the parameters used in the spec, including the paths and queries.
- [.spectral.yaml](./spec/openapi/.spectral.yaml): Contains the rules from [Spectral](https://meta.stoplight.io/docs/spectral/README.md) linter. Reference your preferred rules from [here](https://meta.stoplight.io/docs/spectral/docs/reference/openapi-rules.md).

### Defining Endpoints
All endpoints are defined in the [openapi.yaml](./spec/openapi/openapi.yaml) file. The endpoints are flushed out in the resources (i.e. ./rsources/pets) file found in [resources](./spec/openapi/resources). Each endpoint is defined as follows:
```yaml
paths:
  /pets:
    $ref: "./resources/pets.yaml"
  /pets/{petId}:
    $ref: "./resources/pet.yaml"
```

### Detailing Endpoint Methods
Each endpoint method is defined in the resources file (i.e. ./rsources/pets) found in [resources](./spec/openapi/resources). Each method is defined as follows:
```yaml
get:
  summary: Detail
  operationId: showPetById
  description: Info for a specific pet
  tags:
    - pets
  parameters:
    - $ref: "../parameters/path/petId.yaml"
  responses:
    '200':
      description: Expected response to a valid request
      content:
        application/json:
          schema:
            $ref: "../schemas/Pet.yaml"
    default:
      $ref: "../responses/UnexpectedError.yaml"
```

### Defining Schemas
All schemas are specified in the resources file and defined in their respective schemas file. After a schema file is created add the schema file to the schema index file [_index.yaml](./spec/openapi/schemas/_index.yaml). The schema index file is used to generate the [openapi-v1.0.yaml](./spec/openapi/openapi-v1.0.yaml) file.
Below is an example of a schema:
```shell
type: object
required:
  - id
  - name
properties:
  id:
    type: integer
    format: int64
  name:
    type: string
  tag:
    type: string
```

Then we go into the _index.yaml file and add the schema file to the list of schemas:
```shell
- $ref: "./Pet.yaml"
```

### Defining Responses
All responses are specified in the resources file and defined in their respective responses file. After a response file is created add the response file to the response index file [_index.yaml](./spec/openapi/responses/_index.yaml). The response index file is used to generate the [openapi-v1.0.yaml](./spec/openapi/openapi-v1.0.yaml) file.
Below is an example of a response:
```shell
description: Unexpected error
content:
  application/json:
    schema:
      $ref: "../schemas/Error.yaml"
```

Then we go into the _index.yaml file and add the response file to the list of responses:
```shell
- $ref: "./UnexpectedError.yaml"
```

### Defining Parameters
All parameters are specified in the resources file and defined in their respective parameters file. After a parameter file is created add the parameter file to the parameter index file [_index.yaml](./spec/openapi/parameters/_index.yaml). The parameter index file is used to generate the [openapi-v1.0.yaml](./spec/openapi/openapi-v1.0.yaml) file.
Below is an example of a parameter:
```shell
name: petId
in: path
description: ID of pet to use
required: true
schema:
  type: integer
  format: int64
```

Then we go into the _index.yaml file and add the parameter file to the list of parameters:
```shell
- $ref: "./path/petId.yaml"
```

## Implementing the Services
The generated code should never be edited by you directly, instead you should write all your code in the [/src](./src) folder. 

### Adding Business Logic
The business logic should be written in the [./src/service](./src/service) folder. 
The service folder is organized by tags defined previously in the spec. One of the tags was pet, so there is a service called `pet_service.go` in the [./src/service/](./src/service/) folder.

Below is an example of implementing the service that was autogenerated earlier off from OpenAPI spec, here only one of the services is implemented ShowPetById:
```go
// MyPetApiService PetApiServicer is autogenerated code found in the build directory
// Declares a new type, MyPersonApiService, which is a struct with no fields.
// It will be used to implement the openapi.PersonAPIServicer interface.
type MyPetApiService struct{}

// NewMyPetApiService This is a constructor function for MyPersonApiService.
// It returns a new instance of MyPersonApiService, but the return type is openapi.PersonAPIServicer, which is an interface type.
// This means the function promises to return something that satisfies the PersonAPIServicer interface, which in this case is a *MyPersonApiService.
func NewMyPetApiService() openapi.PetsAPIServicer {
	return &MyPetApiService{}
}

// CreatePets is a method on MyPersonApiService that satisfies the PersonAPIServicer interface.
func (m MyPetApiService) CreatePets(ctx context.Context) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

// ListPets is a method on MyPersonApiService that satisfies the PersonAPIServicer interface.
func (m MyPetApiService) ListPets(ctx context.Context, i int32) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

// ShowPetById is a method on MyPersonApiService that satisfies the PersonAPIServicer interface.
func (m MyPetApiService) ShowPetById(ctx context.Context, s string) (openapi.ImplResponse, error) {
	pet := openapi.Pet{Id: 23, Name: "Dog", Tag: "hello"}

	logger.Logger.Info("Person value is called", zap.String("pet", pet.Name))

	return openapi.Response(http.StatusOK, pet), nil
}
```
### Connecting the Service to the Controller
The step above just added the logic, now we want to serve this business logic. So we register our service to our controller, we do this in the [./src/controller](./src/controller) folder.
```go
func SetupServicesAndControllers() (chi.Router, error) {
	// Create services
	personApiService := service.NewMyPersonApiService()
	newMyPetApiService := service.NewMyPetApiService()

	// Set up router
	router := openapi.NewRouter(petsApiController)

	return router, nil
}
```
Congratulations! You implemented your first OpenAPI Spec First API! Now you can run the server and test it out.


